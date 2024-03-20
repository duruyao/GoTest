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

type Option struct {
	xName           string
	yName           string
	title           string
	link            string
	subtitle        string
	subLink         string
	lineChartSymbol string
	lineChartColor  string
}

func (o *Option) XName() string {
	return o.xName
}

func (o *Option) YName() string {
	return o.yName
}

func (o *Option) Title() string {
	return o.title
}

func (o *Option) Link() string {
	return o.link
}

func (o *Option) Subtitle() string {
	return o.subtitle
}

func (o *Option) SubLink() string {
	return o.subLink
}

func (o *Option) LineChartSymbol() string {
	return o.lineChartSymbol
}

func (o *Option) LineChartColor() string {
	return o.lineChartColor
}
