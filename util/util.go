package util

import (
	"bytes"
	"encoding/base32"
	"log"
	"strconv"
	"strings"
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

func StringToFloat64Must(s string) float64 {
	if len(s) == 0 {
		return 0
	}
	n, e := strconv.ParseFloat(s, 64)
	if e != nil {
		log.Fatalln(e)
	}
	return n
}

func StringsToJsArray(s []string) string {
	if len(s) == 0 {
		return `[]`
	}
	return `["` + strings.Join(s, `", "`) + `"]`
}
