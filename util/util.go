package util

import (
	"bytes"
	"encoding/base32"
	"log"
	"math"
	"strconv"
	"strings"
	"text/template"
)

func TemplateToStringMust(text string, data any) string {
	name := base32.StdEncoding.EncodeToString([]byte(text))
	t := template.Must(template.New(name).Parse(text))
	var b bytes.Buffer
	if e := t.Execute(&b, data); e != nil {
		log.Fatalln(e)
	}
	return b.String()
}

func StringToFloatMust(s string) float64 {
	n, e := strconv.ParseFloat(s, 64)
	if e != nil {
		log.Fatalln(e)
	}
	return n
}

func ChangeFloatPrecision(f float64, prec int) float64 {
	factor := math.Pow(10, float64(prec))
	return math.Round(f*factor) / factor
}

func StringsToJsArray(s []string) string {
	if len(s) == 0 {
		return `[]`
	}
	return `["` + strings.Join(s, `", "`) + `"]`
}

func RemoveExt(path string) string {
	for i := len(path) - 1; i >= 0 && path[i] != '/'; i-- {
		if path[i] == '.' {
			return path[:i]
		}
	}
	return path
}
