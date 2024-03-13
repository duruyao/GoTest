package graph

import (
	"github.com/duruyao/gotest/accuracy"
	"github.com/duruyao/gotest/util"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/event"
	"github.com/go-echarts/go-echarts/v2/opts"
	"io"
)

type Line struct{}

func (Line) RenderAccuracy(r io.Writer, results []accuracy.Result) error {
	xAxis := make([]string, 0)
	links := make([]string, 0)
	series := make([]opts.LineData, 0)
	for i := range results {
		xAxis = append(xAxis, results[i].DateMust())
		links = append(links, results[i].HtmlPageRecordUrl())
		series = append(series, opts.LineData{Value: util.StringToFloat64Must(results[i].Value)})
	}

	line := charts.NewLine()
	JFunc := `function(params) {
		const links = ` + util.StringsToJsArray(links) + `;
		window.open(links[params.dataIndex], '_blank');
	}`
	line.SetGlobalOptions(
		charts.WithTitleOpts(
			opts.Title{
				Title:    results[len(results)-1].TestCaseShortTitle(),
				Link:     results[len(results)-1].TestCasePackageUrl(),
				Subtitle: results[len(results)-1].TestCaseLongTitle(),
				SubLink:  results[len(results)-1].HtmlDirUrl(),
			},
		),
		charts.WithTooltipOpts(opts.Tooltip{Show: opts.Bool(true), Trigger: "axis", Enterable: opts.Bool(true)}),
		charts.WithInitializationOpts(opts.Initialization{Width: "2000px", Height: "800px"}),
		//charts.WithXAxisOpts(opts.XAxis{Name: "Date"}),
		//charts.WithYAxisOpts(opts.YAxis{Name: "Accuracy"}),
		charts.WithEventListeners(event.Listener{EventName: "click", Handler: opts.FuncOpts(JFunc)}),
	)

	line.SetXAxis(xAxis)
	line.AddSeries(results[len(results)-1].TestCaseShortTitle(), series)
	line.SetSeriesOptions(
		charts.WithLineChartOpts(
			opts.LineChart{Smooth: opts.Bool(true), ShowSymbol: opts.Bool(true), SymbolSize: 10, Symbol: "circle"},
		),
	)

	page := components.NewPage()
	page.AddCharts(line)

	return page.Render(io.MultiWriter(r))
}
