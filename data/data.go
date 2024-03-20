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

package data

type Data struct {
	x    []string
	y    []float64
	link []string
}

func (d *Data) Append(x string, y float64, link string) {
	d.x = append(d.x, x)
	d.y = append(d.y, y)
	d.link = append(d.link, link)
}

func (d *Data) N() int {
	return len(d.x)
}

func (d *Data) X(i int) string {
	return d.x[i]
}

func (d *Data) Y(i int) float64 {
	return d.y[i]
}

func (d *Data) Link(i int) string {
	return d.link[i]
}
