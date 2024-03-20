// Copyright 2023-2033 Ryan Du <duruyao@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package data

import (
	"fmt"
	"github.com/duruyao/gotest/conf"
	"github.com/duruyao/gotest/util"
	"github.com/gocarina/gocsv"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

type record struct {
	Id             string `csv:"DIV_ID"`
	Accuracy       string `csv:"ACCURACY"`
	Similarity     string `csv:"COSINE_SIMILARITY"`
	PackagePath    string `csv:"TEST_CASE_PATH"`
	NpuModelPath   string `csv:"NPUMODEL"`
	JobUrl         string `csv:"CI_JOB_URL"`
	ProjectUrl     string `csv:"CI_PROJECT_URL"`
	PipelineUrl    string `csv:"CI_PIPELINE_URL"`
	CommitBranch   string `csv:"CI_COMMIT_BRANCH"`
	CommitShortSha string `csv:"CI_COMMIT_SHORT_SHA"`
	CsvPath        string
}

func (r *record) project() string {
	return path.Base(r.ProjectUrl)
}

func (r *record) csvDir() string {
	return path.Dir(r.CsvPath)
}

func (r *record) htmlPath() string {
	return strings.ReplaceAll(r.CsvPath, "csv", "html")
}

func (r *record) htmlDir() string {
	return path.Dir(r.htmlPath())
}

func (r *record) dateMust() string {
	s := path.Base(r.CsvPath)
	s = strings.ReplaceAll(s, ".csv", "")
	s = strings.ReplaceAll(s, "-"+r.CommitShortSha, "")
	s = strings.ReplaceAll(s, r.project()+"-result-", "")
	t, e := time.Parse(time.RFC3339, s)
	if e != nil {
		log.Fatalln(e)
	}
	return t.Format("06-01-02") // YY-MM-DD
}

func (r *record) accuracyMust() float64 {
	if len(r.Accuracy) > 0 {
		return util.StringToFloatMust(r.Accuracy)
	}
	return 0.0
}

func (r *record) similarityMust() float64 {
	if len(r.Similarity) > 2 {
		s, sum := strings.Split(r.Similarity[1:len(r.Similarity)-1], ", "), 0.0
		for i := range s {
			sum += util.StringToFloatMust(s[i])
		}
		return util.ChangeFloatPrecision(sum/float64(len(s)), 3)
	}
	return 0.0
}

func (r *record) recordUrl() string {
	return fmt.Sprintf("%s%s#%s", conf.FileServerAddr, r.htmlPath(), r.Id)
}

func (r *record) packageTitleMust() string {
	if model := path.Base(r.NpuModelPath); strings.Contains(model, "--") {
		return strings.Join(strings.Split(model, "--")[:3], "-")
	}
	s := strings.Split(util.RemoveExt(r.NpuModelPath), "/")
	return s[len(s)-1] + "-" + s[len(s)-2]
}

func (r *record) packageUrlMust() string {
	s := strings.Split(r.PackagePath, r.project()+"/")
	return fmt.Sprintf("%s/artifacts/raw/%s?inline=false", r.JobUrl, s[len(s)-1])
}

func (r *record) htmlDirTitle() string {
	return fmt.Sprintf("ai_npu/%s/%s", r.project(), r.CommitBranch)
}

func (r *record) htmlDirUrl() string {
	return fmt.Sprintf("%s%s", conf.FileServerAddr, r.htmlDir())
}

func queryRecord(file, id string) (*record, error) {
	f, e := os.Open(file)
	if e != nil {
		return nil, e
	}
	defer func() {
		if e = f.Close(); e != nil {
			log.Fatalln(e)
		}
	}()
	var records []record
	if e := gocsv.UnmarshalFile(f, &records); e != nil {
		return nil, e
	}
	for _, r := range records {
		if r.Id == id {
			r.CsvPath = file
			return &r, nil
		}
	}
	return nil, nil
}
