package implements

type Square struct {
	Side float64
}

func (s Square) CalArea() float64 {
	return s.Side * s.Side
}

func (s Square) CalPerimeter() float64 {
	return 4 * s.Side
}
