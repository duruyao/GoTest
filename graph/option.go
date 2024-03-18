package graph

type Option interface {
	XName() string
	YName() string
	Title() string
	Link() string
	Subtitle() string
	SubLink() string
	LineChartSymbol() string
	LineChartColor() string
}
