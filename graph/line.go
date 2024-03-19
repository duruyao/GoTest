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
	for i := 0; i < data.N(); i++ {
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
			},
		),
		charts.WithTooltipOpts(opts.Tooltip{Show: opts.Bool(true), Trigger: "axis", Enterable: opts.Bool(true)}),
		charts.WithInitializationOpts(opts.Initialization{Width: "2000px", Height: "800px"}),
		charts.WithXAxisOpts(opts.XAxis{Name: option.XName()}),
		charts.WithYAxisOpts(opts.YAxis{Name: option.YName()}),
		charts.WithEventListeners(event.Listener{EventName: "click", Handler: opts.FuncOpts(JFunc)}),
	)

	line.SetXAxis(xAxis)
	line.AddSeries(option.Title(), series)
	line.SetSeriesOptions(
		charts.WithLineChartOpts(
			opts.LineChart{
				Smooth:     opts.Bool(false),
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
			opts.MarkPointStyle{SymbolSize: 80, Label: &opts.Label{Show: opts.Bool(true)}}),
	)

	page := components.NewPage()
	page.SetPageTitle(conf.App)
	page.AddCharts(line)

	return page.Render(io.MultiWriter(r))
}
