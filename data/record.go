package data

import (
	"fmt"
	"github.com/duruyao/gotest/conf"
	"github.com/duruyao/gotest/util"
	"github.com/gocarina/gocsv"
	"log"
	"math/rand"
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
	return util.StringToFloat64Must(r.Accuracy)
}

func (r *record) similarityMust() float64 {
	// TODO: finish
	return float64(rand.Int() % 100 / 100)
}

func (r *record) recordUrl() string {
	return fmt.Sprintf("%s%s#%s", conf.FileServerAddr, r.htmlPath(), r.Id)
}

func (r *record) packageTitleMust() string {
	return strings.Join(strings.Split(path.Base(r.NpuModelPath), "--")[:3], "-")
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
