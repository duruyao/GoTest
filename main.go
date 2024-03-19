package main

import (
	"fmt"
	"github.com/duruyao/gotest/args"
	"github.com/duruyao/gotest/conf"
	"github.com/duruyao/gotest/data"
	"github.com/duruyao/gotest/graph"
	"github.com/duruyao/gotest/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	if args.WantHelp() {
		fmt.Println(args.Usage())
		return
	}
	if args.WantVersion() {
		fmt.Println(conf.VersionSerial())
		return
	}

	router := gin.Default()
	router.GET("/history", func(c *gin.Context) {
		params := struct {
			Project        string `form:"project" binding:"required,oneof=vc0728 vc0768"`
			TestType       string `form:"test_type" binding:"required,oneof=accuracy similarity"`
			Branch         string `form:"branch" binding:"required"`
			TestCaseId     string `form:"test_case_id" binding:"required"`
			CommitShortSha string `form:"commit_short_sha"`
		}{}

		if err := c.ShouldBindQuery(&params); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		testResultDir := util.TemplateToStringMust(conf.CsvResultDirTmpl, params)

		if history, err := data.QueryHistory(testResultDir, params.TestCaseId, params.CommitShortSha); err != nil {
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
	log.Fatalln(router.Run(args.Host()))
}
