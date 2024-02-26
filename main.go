package main

import (
	"github.com/duruyao/gotest/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const fileServerAddress = "http://10.0.13.134:3927"
const testResultsDirFmt = `/opt/share0/gitlab-ci/{{.Project}}/test-result/{{.TestStages}}/{{.Branch}}/{{.FileType}}`

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
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "valid test_type value: " + testType})
			return
		}

		branch := c.DefaultQuery("branch", "dev")

		testCaseId := c.Query("test_case_id")

		testResultsUrl := fileServerAddress + util.TemplateToStringMust(testResultsDirFmt, struct {
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
		c.IndentedJSON(http.StatusOK, gin.H{
			"project":          project,
			"test_type":        testType,
			"teat_stage":       testStages,
			"branch":           branch,
			"test_case_id":     testCaseId,
			"test_results_url": testResultsUrl,
		})
	})
	log.Fatalln(router.Run())
}
