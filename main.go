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
		param := struct {
			Branch         string `form:"branch" binding:"required"`
			Project        string `form:"project" binding:"required,oneof=vc0728 vc0768"`
			TestType       string `form:"test_type" binding:"required,oneof=accuracy similarity"`
			TestCaseId     string `form:"test_case_id" binding:"required"`
			CommitShortSha string `form:"commit_short_sha"`
		}{}

		if err := c.ShouldBindQuery(&param); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		testResultDir := util.TemplateToStringMust(conf.CsvResultDirFmt, param)

		if history, err := data.QueryHistory(testResultDir, param.TestCaseId, param.CommitShortSha); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		} else if history.Data.N() == 0 {
			c.IndentedJSON(http.StatusNotFound, gin.H{"history": history})
			return
		} else {
			if err := graph.Render(c.Writer, &history.Data, &history.Option); err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
			}
		}
	})

	log.Fatalln(router.Run(conf.ListeningAddr))
}
