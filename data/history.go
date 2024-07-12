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
	"io/fs"
	"path/filepath"
	"sort"
	"strings"
)

type History struct {
	Data   Data
	Option Option
}

const (
	AccuracyTest   = "accuracy"
	SimilarityTest = "similarity"
)

func QueryHistory(dir, id, testType, lastCommit string, crossPlatform bool) (*History, error) {
	var csvFiles []string
	if e := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".csv") {
			csvFiles = append(csvFiles, dir+"/"+info.Name())
		}
		return nil
	}); e != nil {
		return nil, e
	}
	sort.Strings(csvFiles)

	end := len(csvFiles)
	if len(lastCommit) > 0 {
		for i := len(csvFiles) - 1; i >= 0; i-- {
			if strings.Contains(csvFiles[i], lastCommit) {
				end = i + 1
				break
			}
		}
	}

	r, e := (*record)(nil), error(nil)
	if testType == AccuracyTest {
		history := &History{}
		for i := range csvFiles[:end] {
			if r, e = queryRecord(csvFiles[i], id); r != nil && e == nil {
				history.Data.Append(r.dateMust(), r.accuracyMust(), r.recordUrl())
			}
		}
		if r != nil {
			history.Option = Option{
				xName:           "Date (YY-MM-DD)",
				yName:           "Accuracy (0 ~ 1)",
				title:           r.packageTitleMust(),
				link:            r.packageUrlMust(),
				subtitle:        r.htmlDirTitle(),
				subLink:         r.htmlDirUrl(),
				lineChartSymbol: "circle",
				lineChartColor:  "",
			}
		}
		return history, nil
	} else if testType == SimilarityTest && !crossPlatform {
		history := &History{}
		for i := range csvFiles[:end] {
			if r, e = queryRecord(csvFiles[i], id); r != nil && e == nil {
				history.Data.Append(r.dateMust(), r.similarityMust(), r.recordUrl())
			}
		}
		if r != nil {
			history.Option = Option{
				xName:           "Date (YY-MM-DD)",
				yName:           "Average similarity between Caffe model and NPU model (0 ~ 1)",
				title:           r.packageTitleMust(),
				link:            r.packageUrlMust(),
				subtitle:        r.htmlDirTitle(),
				subLink:         r.htmlDirUrl(),
				lineChartSymbol: "diamond",
				lineChartColor:  "#8E44AD",
			}
		}
		return history, nil
	} else if testType == SimilarityTest && crossPlatform {
		history := &History{}
		for i := range csvFiles[:end] {
			if r, e = queryRecord(csvFiles[i], id); r != nil && e == nil {
				history.Data.Append(r.dateMust(), r.similarity2Must(), r.recordUrl())
			}
		}
		if r != nil {
			history.Option = Option{
				xName:           "Date (YY-MM-DD)",
				yName:           "Average similarity of Caffe models on AMD64 and ARMv7 (0 ~ 1)",
				title:           r.packageTitleMust(),
				link:            r.packageUrlMust(),
				subtitle:        r.htmlDirTitle(),
				subLink:         r.htmlDirUrl(),
				lineChartSymbol: "triangle",
				lineChartColor:  "#CD5C5C",
			}
		}
		return history, nil
	}

	return nil, nil
}
