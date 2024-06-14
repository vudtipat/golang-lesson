package main

import "fmt"

func reference_var(ptr_s *string, newStr string) {
	*ptr_s = newStr
}

func main() {
	s := "Hello"
	fmt.Println(s)
	reference_var(&s, "World")
	fmt.Println(s)
}
