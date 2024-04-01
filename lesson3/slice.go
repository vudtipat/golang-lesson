package main

import (
	"fmt"
)

func main() {

	arr := [10]int {10,20,30,40,50,60,70,80,90,100}

	slice1 := []int {10,20,30}
	fmt.Println(slice1)

	slice := arr[1:8]
	fmt.Println(slice)
	fmt.Println("slice len = ",len(slice))
	fmt.Println("slice cap = ",cap(slice))

	sub_slice := slice[0:3]
	fmt.Println(sub_slice)
	fmt.Println("sub_slice len = ",len(sub_slice))
	fmt.Println("sub_slice cap = ",cap(sub_slice))
	sub_slice = append(sub_slice,0,10)
	fmt.Println(sub_slice)
}