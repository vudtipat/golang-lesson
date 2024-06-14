package main

import "fmt"

type Student struct {
	name   string
	rollNo int
	marks  []int
	grades map[string]int
}

func main() {
	var std Student
	fmt.Printf("%+v\n", std)

	var std1 Student = Student{
		name:   "Tum",
		rollNo: 1,
	}
	fmt.Printf("%+v\n", std1)

	std2 := Student{
		name:   "Tum2",
		rollNo: 1,
	}
	fmt.Printf("%+v\n", std2)

	std2.rollNo = 2
	fmt.Printf("%+v\n", std2)
}
