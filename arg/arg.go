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

package arg

import (
	"bytes"
	"flag"
	"github.com/duruyao/gotest/conf"
	"log"
	"os"
	"sync"
	"text/template"
)

const UsageTmpl = `{{.Logo}}
Usage: {{.Exec}} [OPTIONS]

GoTest parses test results and creates interactive visualizations in the web

Options:
    -h, --help                  Display this help message
    -v, --version               Print version information and quit
    --dir  STRING               Directory of dataset (default: '{{.Dir}}')
    --host STRING               Host address to listen (default: '{{.Host}}')

Examples:
    {{.Exec}}
    {{.Exec}} --dir {{.Dir}}
    {{.Exec}} --dir {{.Dir}} --host {{.Host}} 

See more about {{.App}} at {{.Link}}
`

type args struct {
	dir         string
	host        string
	wantHelp    bool
	wantVersion bool
	parseOnce   sync.Once
}

func (a *args) parse() {
	a.parseOnce.Do(func() {
		flag.Parse()
	})
}

var a args

func Dir() string {
	a.parse()
	return a.dir
}

func Host() string {
	a.parse()
	return a.host
}

func WantHelp() bool {
	a.parse()
	return a.wantHelp
}

func WantVersion() bool {
	a.parse()
	return a.wantVersion
}

func Usage() string {
	a.parse()
	tmpl := template.Must(template.New("usage tmpl").Parse(UsageTmpl))
	data := struct {
		Logo string
		Exec string
		Dir  string
		Host string
		App  string
		Link string
	}{
		Logo: conf.Logo,
		Exec: os.Args[0],
		Dir:  conf.DefaultDir,
		Host: conf.DefaultHost,
		App:  conf.App,
		Link: conf.Link,
	}
	buf := bytes.Buffer{}
	if e := tmpl.Execute(&buf, data); e != nil {
		log.Fatalln(e)
	}
	return buf.String()
}
