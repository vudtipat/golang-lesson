package main

import "fmt"

type Circle struct {
	x      float64
	y      float64
	radius float64
	area   float64
}

func calAreaCircle(c Circle) float64 {
	pi := 3.14
	area := pi * c.radius * c.radius
	return area
}

func calAreaCircleRefArea(c *Circle) {
	pi := 3.14
	area := pi * c.radius * c.radius
	(*c).area = area
}

func main() {
	c1 := Circle{
		x:      5,
		y:      5,
		radius: 5,
	}

	fmt.Printf("%+v\n", c1)
	area := calAreaCircle(c1)
	fmt.Println(area)

	calAreaCircleRefArea(&c1)
	fmt.Printf("%+v\n", c1)
}
