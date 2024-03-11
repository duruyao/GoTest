package graph

import (
	"github.com/duruyao/gotest/accuracy"
	"github.com/duruyao/gotest/util"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"io"
)

func lineSymbols(results []accuracy.Result) *charts.Line {
	xData, yData := make([]string, 0), make([]float64, 0)
	for _, result := range results {
		xData = append(xData, result.DateMust())
		yData = append(yData, util.StringToFloat64Must(result.Value))
	}
	series := make([]opts.LineData, 0)
	for _, data := range yData {
		series = append(series, opts.LineData{Value: data})
	}

	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    results[len(results)-1].TestCaseName(),
			Link:     results[len(results)-1].TestCaseUrl(),
			Subtitle: results[len(results)-1].SubTitle(),
			SubLink:  results[len(results)-1].HtmlDirUrl(),
		}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true, Trigger: "axis"}),
		charts.WithInitializationOpts(
			opts.Initialization{
				Width:  "2000px",
				Height: "800px",
			}),
	)

	// Put data into instance
	line.SetXAxis(xData).
		AddSeries(results[len(results)-1].TestCaseName(), series).
		SetSeriesOptions(charts.WithLineChartOpts(
			opts.LineChart{Smooth: true, ShowSymbol: true, SymbolSize: 15, Symbol: "diamond"},
		))

	return line
}

type Line struct{}

func (Line) Render(r io.Writer, results []accuracy.Result) error {
	page := components.NewPage()
	page.AddCharts(
		lineSymbols(results),
	)
	return page.Render(io.MultiWriter(r))
}
