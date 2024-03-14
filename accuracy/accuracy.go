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

func (r *Result) TestCasePackageUrl() string {
	s := strings.Split(r.CasePath, r.ProjectNameMust()+"/")
	return fmt.Sprintf("%s/artifacts/raw/%s?inline=false", r.JobUrl, s[len(s)-1])
}

func (r *Result) TestCaseShortTitle() string {
	return strings.Join(strings.Split(path.Base(r.NpuModel), "--")[:3], "-")
}

func (r *Result) HtmlDirUrl() string {
	return fmt.Sprintf("%s%s", conf.FileServerAddress, strings.ReplaceAll(r.CsvDir, "csv", "html"))
}

func (r *Result) TestCaseLongTitle() string {
	return fmt.Sprintf("Accuracy test results for project %s (%s branch)", r.ProjectNameMust(), r.CommitBranch)
}

func (r *Result) HtmlPageRecordUrl() string {
	return fmt.Sprintf("%s%s#%s", conf.FileServerAddress, strings.ReplaceAll(r.CsvPath, "csv", "html"), r.Id)
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

func QueryResults(dir, id, commit string) ([]Result, error) {
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

	idx := len(csvFiles)
	if found := false; len(commit) > 0 {
		for idx, found = 0, false; idx < len(csvFiles); idx++ {
			if contains := strings.Contains(csvFiles[idx], commit); found && (!contains) {
				break
			} else {
				found = contains
			}
		}
	}
	for _, csvFile := range csvFiles[:idx] {
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
