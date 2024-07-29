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

package conf

import (
	"bytes"
	"github.com/duruyao/gotest/util"
	"log"
	"text/template"
)

var DefaultDir = util.GetWorkDirMust() + "/dataset"

const (
	DefaultHost      = "0.0.0.0:4936"
	DirUrlPrefix     = "/dataset"
	CsvResultDirTmpl = `{{.Project}}/test-result/{{.TestType}}/{{.Branch}}/csv`
)

const (
	App               = `GoTest`
	Link              = `https://github.com/duruyao/gotest`
	Version           = `1.0.1`
	ReleaseDate       = `2024-07-29`
	VersionSerialTmpl = `{{.App}} {{.Version}} ({{.ReleaseDate}})`
	Logo              = `
   _____    _______        _
  / ____|  |__   __|      | |
 | |  __  ___ | | ___  ___| |_
 | | |_ |/ _ \| |/ _ \/ __| __|
 | |__| | (_) | |  __/\__ \ |_
  \_____|\___/|_|\___||___/\__|

`
)

func VersionSerial() string {
	tmpl := template.Must(template.New("version serial tmpl").Parse(VersionSerialTmpl))
	data := struct {
		App         string
		Version     string
		ReleaseDate string
	}{
		App:         App,
		Version:     Version,
		ReleaseDate: ReleaseDate,
	}
	buf := bytes.Buffer{}
	if e := tmpl.Execute(&buf, data); e != nil {
		log.Fatalln(e)
	}
	return buf.String()
}
