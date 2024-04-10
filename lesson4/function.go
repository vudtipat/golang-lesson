package main

import (
	"fmt"
)

// return 1 value
func addValue(m int,n int) int {
	return m+n
}

// don't return
func printValue(value string) {
	fmt.Println(value)
}

// multiple return
func findEvenNumber(value []int) (count int, lst []int) {
	count_lst := 0
	var new_lst = []int {}
	for _,data := range value {
		if(data % 2 == 0) {
			count_lst+=1
			new_lst = append(new_lst,data)
		}
	}
	return count_lst,new_lst
}

// variadic function
func sumaryValue(numbers ...int)int {
	sum := 0
	for _,data := range numbers {
		sum += data
	}
	return sum
}

func calArea(r float64) float64{
	return 3.14 * r * r
}

func calPerimeter(r float64) float64{
	return 3.14 * 2 * r
}

func calDiameter(r float64) float64{
	return 2 * r
}

func printResult(radius float64,calfunction func(r float64) float64) {
	result := calfunction(radius)
	fmt.Println("Resutl:",result)
	fmt.Println("Thank you!")
}

func getFunction(query int) func(r float64) float64 {
	query_to_function := map[int]func(r float64) float64{
		1: calArea,
		2: calPerimeter,
		3: calDiameter,
	}
	return query_to_function[query]
}

func main() {
	m := addValue(3,4)
	fmt.Println("m = ",m)

	printValue("Hello world!!")

	count,lst := findEvenNumber([]int {1,2,3,4,5,6,7,8,9,10})
	fmt.Println("Count = ",count,", even is",lst)

	x := sumaryValue(1,2,3,4,5,6,7,8,10)
	fmt.Println("x =",x)

	var query int
	var radius float64

	fmt.Print("Enter the radius of circle: ")
	fmt.Scanf("%f",&radius)
	fmt.Print("Enter \n 1 - area \n 2 - perimeter \n 3 - diameter \nEnter your choice : ")
	fmt.Scanf("%d",&query)

	printResult(radius,getFunction(query))

}