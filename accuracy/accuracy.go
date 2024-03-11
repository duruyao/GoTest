package accuracy

import (
	"fmt"
	"github.com/duruyao/gotest/conf"
	"github.com/gocarina/gocsv"
	"io/fs"
	"log"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type record struct {
	Id             string `csv:"DIV_ID"`
	Value          string `csv:"ACCURACY"`
	CasePath       string `csv:"TEST_CASE_PATH"`
	ProjectUrl     string `csv:"CI_PROJECT_URL"`
	NpuModel       string `csv:"NPUMODEL"`
	JobUrl         string `csv:"CI_JOB_URL"`
	PipelineUrl    string `csv:"CI_PIPELINE_URL"`
	CommitBranch   string `csv:"CI_COMMIT_BRANCH"`
	CommitShortSha string `csv:"CI_COMMIT_SHORT_SHA"`
}

type Result struct {
	CsvDir  string
	CsvPath string
	record
}

func (r *Result) ProjectNameMust() string {
	u, e := url.Parse(r.ProjectUrl)
	if e != nil {
		log.Fatalln(e)
	}
	return path.Base(u.Path)
}

func (r *Result) DateMust() string {
	s := path.Base(r.CsvPath)
	s = strings.ReplaceAll(s, ".csv", "")
	s = strings.ReplaceAll(s, "-"+r.CommitShortSha, "")
	s = strings.ReplaceAll(s, r.ProjectNameMust()+"-result-", "")
	t, e := time.Parse(time.RFC3339, s)
	if e != nil {
		log.Fatalln(e)
	}
	return t.Format("06-01-02") // YY-MM-DD
}

func (r *Result) TestCaseUrl() string {
	return fmt.Sprintf("%s/artifacts/raw/%s?inline=false", r.JobUrl, r.CasePath)
}

func (r *Result) TestCaseName() string {
	t := path.Base(r.NpuModel)
	t = strings.ReplaceAll(t, "--", "-")
	t = strings.ReplaceAll(t, "-.npumodel", "")
	t = strings.ReplaceAll(t, ".npumodel", "")
	return t
}

func (r *Result) HtmlDirUrl() string {
	return fmt.Sprintf("%s%s", conf.FileServerAddress, strings.ReplaceAll(r.CsvDir, "csv", "html"))
}

func (r *Result) SubTitle() string {
	return fmt.Sprintf("Accuracy test results for project %s (%s branch)", r.ProjectNameMust(), r.CommitBranch)
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
			return &r, nil
		}
	}
	return nil, nil
}

func QueryResults(dir, id string) ([]Result, error) {
	var results []Result

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

	for _, csvFile := range csvFiles {
		if r, e := queryRecord(csvFile, id); r != nil && e == nil {
			results = append(results, Result{
				CsvDir:  dir,
				CsvPath: csvFile,
				record:  *r,
			})
		}
	}

	return results, nil
}