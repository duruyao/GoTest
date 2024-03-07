package data

import (
	"errors"
	"fmt"
	"github.com/gocarina/gocsv"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const fileServerAddress = "http://10.0.13.134:3927"

type record struct {
	Accuracy   string `csv:"ACCURACY"`
	Similarity string `csv:"COSINE_SIMILARITY"`
	CaseId     string `csv:"DIV_ID"`
	CommitId   string `csv:"CI_COMMIT_SHORT_SHA"`
}

type Result struct {
	Date       time.Time
	Value      float64
	HtmlDocUrl string
	record
}

func queryRecordFromFile(filename, targetCaseId string) (record, error) {
	f, e := os.Open(filename)
	if e != nil {
		return record{}, e
	}
	defer func() {
		if e = f.Close(); e != nil {
			log.Fatalln(e)
		}
	}()
	var records []record
	if e := gocsv.UnmarshalFile(f, &records); e != nil {
		return record{}, e
	}
	for _, r := range records {
		if r.CaseId == targetCaseId {
			return r, nil
		}
	}
	return record{}, errors.New(fmt.Sprintf("No such CaseId: %s", targetCaseId))
}

func QueryResultsFromDir(resultsDir, caseId string) ([]Result, error) {
	var results []Result

	var csvFiles []string
	if e := filepath.Walk(resultsDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".csv") {
			csvFiles = append(csvFiles, resultsDir+"/"+info.Name())
		}
		return nil
	}); e != nil {
		return nil, e
	}
	sort.Strings(csvFiles)

	for _, csvFile := range csvFiles {
		if r, e := queryRecordFromFile(csvFile, caseId); e != nil {
			log.Println(e)
		} else {
			results = append(results, Result{
				Date:       time.Time{},
				Value:      0,
				HtmlDocUrl: fileServerAddress + strings.ReplaceAll(csvFile, "csv", "html"),
				record:     r,
			})
		}
	}

	return results, nil
}
