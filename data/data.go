package data

type Data struct {
	x    []string
	y    []float64
	link []string
}

//func (d *Data) Init(len, cap int) {
//	d.x = make([]string, len, cap)
//	d.y = make([]float64, len, cap)
//	d.link = make([]string, len, cap)
//}

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
