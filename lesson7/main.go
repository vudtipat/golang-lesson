package main

import (
	"fmt"
	. "lesson7/implements"
	. "lesson7/interfaces"
)

func printData(s Shape) {
	fmt.Printf("%+v\n", s)
	fmt.Println(s.CalArea())
	fmt.Println(s.CalPerimeter())
}

func main() {
	r := Rect{
		Breadth: 5,
		Length:  3,
	}
	printData(r)
	s := Square{
		Side: 5,
	}
	printData(s)
}
