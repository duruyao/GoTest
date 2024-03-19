package conf

import (
	"bytes"
	"log"
	"text/template"
)

const (
	DefaultHost      = "0.0.0.0:4936"
	FileServerAddr   = "http://10.0.13.134:3927"
	CsvResultDirTmpl = `/opt/gitlab-data/gitlab-test/{{.Project}}/test-result/{{.TestType}}/{{.Branch}}/csv`
)

const (
	App               = `GoTest`
	Link              = `https://github.com/duruyao/gotest`
	Version           = `0.9.8`
	ReleaseDate       = `2024-03-19`
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
