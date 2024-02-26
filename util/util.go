package util

import (
	"bytes"
	"encoding/base32"
	"log"
	"text/template"
)

func TemplateToStringMust(text string, data any) string {
	name := base32.StdEncoding.EncodeToString([]byte(text))
	t := template.Must(template.New(name).Parse(text))
	var b bytes.Buffer
	if err := t.Execute(&b, data); err != nil {
		log.Fatalln(err)
	}
	return b.String()
}
