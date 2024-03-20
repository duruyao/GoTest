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

package main

import (
	"fmt"
	"github.com/duruyao/gotest/arg"
	"github.com/duruyao/gotest/conf"
	"github.com/duruyao/gotest/data"
	"github.com/duruyao/gotest/graph"
	"github.com/duruyao/gotest/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	if arg.WantHelp() {
		fmt.Println(arg.Usage())
		return
	}
	if arg.WantVersion() {
		fmt.Println(conf.VersionSerial())
		return
	}

	router := gin.Default()
	router.GET("/history", func(c *gin.Context) {
		param := struct {
			Project        string `form:"project" binding:"required,oneof=vc0728 vc0768"`
			TestType       string `form:"test_type" binding:"required,oneof=accuracy similarity"`
			Branch         string `form:"branch" binding:"required"`
			TestCaseId     string `form:"test_case_id" binding:"required"`
			CommitShortSha string `form:"commit_short_sha"`
		}{}

		if err := c.ShouldBindQuery(&param); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		testResultDir := util.TemplateToStringMust(conf.CsvResultDirTmpl, param)

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
	log.Fatalln(router.Run(arg.Host()))
}
