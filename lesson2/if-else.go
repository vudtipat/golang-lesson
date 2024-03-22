package main
import (
	"fmt"
)

func main()  {
	x := 10

	if x < 9 {
		fmt.Println("x less than 9")
	}else {
		fmt.Println("x more than 9")
	}

	x = 15
	if x < 9 {
		fmt.Println("x less than 9")
	}else if x >= 10 && x <=15{
		fmt.Println("x between 10 and 15")
	}else {
		fmt.Println("x more than 15")
	}

	
}