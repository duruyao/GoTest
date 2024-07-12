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

package graph

import (
	"github.com/duruyao/gotest/conf"
	"github.com/duruyao/gotest/util"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/event"
	"github.com/go-echarts/go-echarts/v2/opts"
	"io"
)

func Render(r io.Writer, data Data, option Option) error {
	xAxis := make([]string, 0)
	links := make([]string, 0)
	series := make([]opts.LineData, 0)
	k := 0
	for k = 0; k < data.N() && data.Y(k) == 0; k++ { // skip a lot of empty data in the early stage
	}
	for i := k; i < data.N(); i++ {
		xAxis = append(xAxis, data.X(i))
		links = append(links, data.Link(i))
		series = append(series, opts.LineData{Value: data.Y(i)})
	}

	line := charts.NewLine()
	JFunc := `function(params) {
		const links = ` + util.StringsToJsArray(links) + `;
		window.open(links[params.dataIndex], '_blank');
	}`
	line.SetGlobalOptions(
		charts.WithTitleOpts(
			opts.Title{
				Title:    option.Title(),
				Link:     option.Link(),
				Subtitle: option.Subtitle(),
				SubLink:  option.SubLink(),
				Right:    "0",
			},
		),
		charts.WithTooltipOpts(opts.Tooltip{Show: opts.Bool(true), Trigger: "axis", Enterable: opts.Bool(true)}),
		charts.WithInitializationOpts(opts.Initialization{Width: "2000px", Height: "800px"}),
		charts.WithXAxisOpts(
			opts.XAxis{
				Name:         option.XName(),
				NameLocation: "center",
				NameGap:      40,
				AxisLabel:    &opts.AxisLabel{Color: "black"},
			},
		),
		charts.WithYAxisOpts(
			opts.YAxis{
				Name:         option.YName(),
				NameLocation: "center",
				NameGap:      50,
				AxisLabel:    &opts.AxisLabel{Color: "black"},
			},
		),
		charts.WithEventListeners(event.Listener{EventName: "click", Handler: opts.FuncOpts(JFunc)}),
	)

	line.SetXAxis(xAxis)
	line.AddSeries(option.Title(), series)
	line.SetSeriesOptions(
		charts.WithLineChartOpts(
			opts.LineChart{
				Smooth:     opts.Bool(true),
				ShowSymbol: opts.Bool(true),
				SymbolSize: 10,
				Symbol:     option.LineChartSymbol(),
				Color:      option.LineChartColor(),
			},
		),
		charts.WithMarkPointNameTypeItemOpts(
			opts.MarkPointNameTypeItem{Name: "Maximum", Type: "max"},
		),
		charts.WithMarkPointStyleOpts(
			opts.MarkPointStyle{SymbolSize: 70, Label: &opts.Label{Show: opts.Bool(true)}}),
	)

	page := components.NewPage()
	page.SetPageTitle(conf.App)
	page.AddCharts(line)

	return page.Render(io.MultiWriter(r))
}
