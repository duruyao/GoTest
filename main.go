package main

import (
	"github.com/duruyao/gotest/conf"
	"github.com/duruyao/gotest/data"
	"github.com/duruyao/gotest/graph"
	"github.com/duruyao/gotest/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/history", func(c *gin.Context) {
		project := c.DefaultQuery("project", "vc0728")
		if !map[string]bool{"vc0728": true, "vc0768": true}[project] {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "valid project value: " + project})
			return
		}

		testType := c.DefaultQuery("test_type", "accuracy")
		testStages, ok := map[string]string{"accuracy": "convert-infer", "similarity": "convert-dump-compare"}[testType]
		if !ok {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid test_type value: " + testType})
			return
		}

		branch := c.DefaultQuery("branch", "dev")

		testCaseId := c.Query("test_case_id")

		commitShortSha := c.DefaultQuery("commit_short_sha", "")

		testResultsDir := util.TemplateToStringMust(conf.CsvResultDirFmt, struct {
			Project    string
			TestStages string
			Branch     string
		}{
			Project:    project,
			TestStages: testStages,
			Branch:     branch,
		})

		queryFunc := map[string]func(string, string, string) (*data.History, error){
			"accuracy":   data.QueryAccuracyHistory,
			"similarity": data.QuerySimilarityHistory,
		}
		if history, err := queryFunc[testType](testResultsDir, testCaseId, commitShortSha); history.Data.N() == 0 || err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"project":          project,
				"test_type":        testType,
				"test_stages":      testStages,
				"branch":           branch,
				"test_case_id":     testCaseId,
				"commit_short_sha": commitShortSha,
				"test_results_dir": testResultsDir,
				"history":          history,
				"error":            err,
			})
		} else {
			err := graph.Render(c.Writer, &history.Data, &history.Option)
			if err != nil {
				c.IndentedJSON(http.StatusBadRequest, gin.H{
					"project":          project,
					"test_type":        testType,
					"test_stages":      testStages,
					"branch":           branch,
					"test_case_id":     testCaseId,
					"commit_short_sha": commitShortSha,
					"test_results_dir": testResultsDir,
					"history":          history,
					"error":            err,
				})
			}
		}
	})

	log.Fatalln(router.Run(conf.ListeningAddr))
}
