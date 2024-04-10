package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello ")
	defer fmt.Println("!!!")
	fmt.Println("World ")
}