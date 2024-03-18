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
