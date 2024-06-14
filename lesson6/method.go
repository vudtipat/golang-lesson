package main

import "fmt"

type Circle struct {
	x      float32
	y      float32
	radius float32
	area   float32
}

func (c *Circle) calAreaCircle() float32 {
	area := 3.14 * c.radius * c.radius
	c.area = area
	return area
}

func main() {
	cir1 := Circle{
		x:      1,
		y:      2,
		radius: 5,
	}

	fmt.Printf("%+v\n", cir1)
	area := cir1.calAreaCircle()
	fmt.Println(area)
	fmt.Printf("%+v\n", cir1)

}
