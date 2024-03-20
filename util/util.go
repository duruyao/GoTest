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

package util

import (
	"bytes"
	"encoding/base32"
	"log"
	"math"
	"os"
	"path/filepath"
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

func GetWorkDirMust() string {
	d, e := os.Getwd()
	if e != nil {
		log.Fatalln(e)
	}
	return d
}

func RelativePathMust(base, target string) string {
	f, e := filepath.Rel(base, target)
	if e != nil {
		log.Fatalln(e)
	}
	return f
}
