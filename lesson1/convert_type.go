package main
import (
	"fmt"
	"strconv"
)

func main()  {
	var1 := 10
	var2 := 3.14

	var var1_float float32 = float32(var1)
	var var2_float int = int(var2)

	fmt.Printf("%.2f\n",var1_float)
	fmt.Printf("%d",var2_float)


	var s string = "200"
	i,err := strconv.Atoi(s)

	fmt.Printf("%v, %T\n",i,i)
	fmt.Printf("%v, %T\n",err,err)

	j,err := strconv.Atoi("300aa")
	fmt.Printf("%v, %T\n",j,j)
	fmt.Printf("%v, %T\n",err,err)

	k := strconv.Itoa(500)
	fmt.Printf("%q, %T\n",k,k)

	const varx = 30;
	fmt.Println(varx)
}