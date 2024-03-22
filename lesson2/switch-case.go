package main

import (
	"fmt"
)

func main()  {
	i := 800
	switch i {
		case 10:
			fmt.Println("i is 10")
		case 100:
			fmt.Println("i is 100 or 200")
		default:
			fmt.Println("other")
	}

	j := 10
	switch j {
		case 10:
			fmt.Println("i is 10")
			fallthrough
		case 100:
			fmt.Println("i is 100 or 200")
			fallthrough
		default:
			fmt.Println("other")
	}
}