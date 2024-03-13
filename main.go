package main

import (
	"github.com/duruyao/gotest/accuracy"
	"github.com/duruyao/gotest/conf"
	"github.com/duruyao/gotest/graph"
	"github.com/duruyao/gotest/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()

	// http://0.0.0.0:4936/history?project=vc0728&test_type=accuracy&branch=dev&test_case_id=QWxleE5ldG5ldzBPdXRsaWVyX1JlbW92ZU91dGxpZXJfUmVtb3ZlRXVjbGlkZWFu
	// http://0.0.0.0:4936/history?project=vc0728&test_type=similarity&branch=dev&test_case_id=MTYvQllQX0pLWV9tb2RlbF8x
	//
	// http://0.0.0.0:4936/history?project=vc0768&test_type=accuracy&branch=dev&test_case_id=QWxleE5ldG5ldzBPdXRsaWVyX1JlbW92ZU91dGxpZXJfUmVtb3ZlRXVjbGlkZWFu
	// http://0.0.0.0:4936/history?project=vc0728&test_type=similarity&branch=dev&test_case_id=MTYvQllQX0pLWV9tb2RlbF8x_1
	router.GET("/history", func(c *gin.Context) {
		project := c.DefaultQuery("project", "vc0728")
		if !map[string]bool{"vc0728": true, "vc0768": true}[project] {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "valid project value: " + project})
			return
		}

		testType := c.DefaultQuery("test_type", "accuracy")
		testStages, ok := map[string]string{"accuracy": "convert-infer", "similarity": "convert-dump-compare"}[testType]
		if !ok {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "valid test_type value: " + testType})
			return
		}

		branch := c.DefaultQuery("branch", "dev")

		testCaseId := c.Query("test_case_id")

		testResultsDir := util.TemplateToStringMust(conf.TestResultsDirFmt, struct {
			Project    string
			TestStages string
			Branch     string
			FileType   string
		}{
			Project:    project,
			TestStages: testStages,
			Branch:     branch,
			FileType:   "csv",
		})
		if testType != "accuracy" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"project":          project,
				"test_type":        testType,
				"test_stages":      testStages,
				"branch":           branch,
				"test_case_id":     testCaseId,
				"test_results_dir": testResultsDir,
				"error":            "Unsupported testType: " + testType,
			})
			return
		}
		if results, err := accuracy.QueryResults(testResultsDir, testCaseId); len(results) == 0 || err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"project":          project,
				"test_type":        testType,
				"test_stages":      testStages,
				"branch":           branch,
				"test_case_id":     testCaseId,
				"test_results_dir": testResultsDir,
				"error":            err,
				"results":          results,
			})
		} else {
			err := graph.Line{}.RenderAccuracy(c.Writer, results)
			if err != nil {
				c.IndentedJSON(http.StatusBadRequest, gin.H{
					"project":          project,
					"test_type":        testType,
					"test_stages":      testStages,
					"branch":           branch,
					"test_case_id":     testCaseId,
					"test_results_dir": testResultsDir,
					"error":            err,
					"results":          results,
				})
			}
		}
	})
	log.Fatalln(router.Run(conf.ListeningAddress))
}
