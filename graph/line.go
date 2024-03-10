package graph

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"io"
)

func lineSymbols(xData interface{}, yData []float64) *charts.Line {
	series := make([]opts.LineData, 0)
	for _, data := range yData {
		series = append(series, opts.LineData{Value: data})
	}

	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Title",
			Subtitle: "Subtitle",
		}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true, Trigger: "axis"}),
	)

	// Put data into instance
	line.SetXAxis(xData).
		AddSeries("Category B", series).
		SetSeriesOptions(charts.WithLineChartOpts(
			opts.LineChart{Smooth: true, ShowSymbol: true, SymbolSize: 15, Symbol: "diamond"},
		))

	return line
}

type Line struct{}

func (Line) Render(r io.Writer, xData interface{}, yData []float64) error {
	page := components.NewPage()
	page.AddCharts(
		lineSymbols(xData, yData),
	)
	return page.Render(io.MultiWriter(r))
}
