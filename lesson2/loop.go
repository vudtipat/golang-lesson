package main
import (
	"fmt"
)

func main(){
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("Up step by 5")

	for i := 0; i < 20; i+=5 {
		fmt.Println(i)
	}

	fmt.Println("alternative for loop")

	j := 10
	for j <=15 {
		fmt.Println(j)
		j++
	}

	fmt.Println("for loop with break")

	for i:=1;i<100;i++ {
		fmt.Println(i)
		if(i % 5 == 0) {
			break
		}
	}

	fmt.Println("for loop with continue")

	for i:=1;i<10;i++ {
		fmt.Println(i)
		if(i % 5 == 0) {
			continue
		}
	}
}