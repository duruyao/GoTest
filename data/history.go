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
	AccuracyTestType   = "accuracy"
	SimilarityTestType = "similarity"
)

func QueryHistory(dir, id, commit string) (*History, error) {
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
	if len(commit) > 0 {
		for i := len(csvFiles) - 1; i >= 0; i-- {
			if strings.Contains(csvFiles[i], commit) {
				end = i + 1
				break
			}
		}
	}

	r, e := (*record)(nil), error(nil)
	if strings.Contains(dir, AccuracyTestType) {
		history := &History{}
		for i := range csvFiles[:end] {
			if r, e = queryRecord(csvFiles[i], id); r != nil && e == nil {
				history.Data.Append(r.dateMust(), r.accuracyMust(), r.recordUrl())
			}
		}
		if r != nil {
			history.Option = Option{
				xName:           "Date",
				yName:           "Accuracy",
				title:           r.packageTitleMust(),
				link:            r.packageUrlMust(),
				subtitle:        r.htmlDirTitle(),
				subLink:         r.htmlDirUrl(),
				lineChartSymbol: "circle",
				lineChartColor:  "",
			}
		}
		return history, nil
	} else if strings.Contains(dir, SimilarityTestType) {
		history := &History{}
		for i := range csvFiles[:end] {
			if r, e = queryRecord(csvFiles[i], id); r != nil && e == nil {
				history.Data.Append(r.dateMust(), r.similarityMust(), r.recordUrl())
			}
		}
		if r != nil {
			history.Option = Option{
				xName:           "Date",
				yName:           "Average Similarity",
				title:           r.packageTitleMust(),
				link:            r.packageUrlMust(),
				subtitle:        r.htmlDirTitle(),
				subLink:         r.htmlDirUrl(),
				lineChartSymbol: "diamond",
				lineChartColor:  "#8E44AD",
			}
		}
		return history, nil
	}

	return nil, nil
}
