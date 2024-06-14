package main

import "fmt"

func main() {
	i := 10
	fmt.Printf("%T %v \n", &i, &i)
	fmt.Printf("%T %v \n", *(&i), *(&i))

	// var ptr_i *int = &i
	// fmt.Println(ptr_i)

	// var ptr_i = &i
	// fmt.Println(ptr_i)

	ptr_i := &i
	fmt.Println(ptr_i)

	s := "Hello"
	fmt.Printf("%T %v \n", s, s)
	ptr_s := &s
	*ptr_s = "world"
	fmt.Printf("%T %v \n", s, s)

}
