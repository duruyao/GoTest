package args

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
    --host STRING               Host address to listen (default: '{{.Host}}')

Examples:
    {{.Exec}}
    {{.Exec}} --host {{.Host}}

See more about {{.App}} at {{.Link}}
`

type args struct {
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
		Host string
		App  string
		Link string
	}{
		Logo: conf.Logo,
		Exec: os.Args[0],
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
