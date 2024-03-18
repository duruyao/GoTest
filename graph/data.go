package graph

type Data interface {
	N() int
	X(int) string
	Y(int) float64
	Link(int) string
}
