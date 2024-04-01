package main
import (
	"fmt"
)

func main(){
	var grades [5]int
	fmt.Println(grades)

	var nGrade [7]int = [7]int{0,1,2,3,4} // or nGrade := [5]int{0,1,2,3,4}
	fmt.Println(nGrade)

	str_grade := [...]string{"A","B","C"}
	fmt.Println(str_grade)
	fmt.Println(len(str_grade))
	fmt.Println(str_grade[0])
	str_grade[1]="A+"
	fmt.Println(str_grade)

	for v:=0;v<len(nGrade);v++ {
		fmt.Println(nGrade[v])
	}

	for idx,element := range nGrade {
		fmt.Println(idx," => ",element)
	}

	narr := [3][2]int{{1,2},{3,4},{5,6}}
	fmt.Println(narr[0][1])
}
