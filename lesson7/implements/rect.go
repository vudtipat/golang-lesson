package implements

type Rect struct {
	Length, Breadth float64
}

func (r Rect) CalArea() float64 {
	return r.Breadth * r.Length
}

func (r Rect) CalPerimeter() float64 {
	return (2 * r.Breadth) + (2 * r.Length)
}
