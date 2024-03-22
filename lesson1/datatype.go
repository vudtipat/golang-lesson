package main
import (
	"fmt"
	"reflect"
)

var global string = "This is global variable"

func main()  {
	// declare type
	var str1 string = "Hello"
	// without declare data type but must assign value and must only use on declare scope
	str2 := "World"

	var balance float32 = 16.95

	var boolean bool = true

	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(balance)
	fmt.Println(boolean)
	fmt.Println(global)
	fmt.Printf("The value is %v and type is %T.\n",str1,str1)
	fmt.Printf("The value is %v and type is %T.\n",balance,balance)
	fmt.Printf("The value is %v and type is %T.\n",boolean,boolean)

	// use reflect package to get type of variable
	fmt.Printf("The value is %v and type is %v.\n",str1,reflect.TypeOf(str1))
	fmt.Printf("The value is %v and type is %v.\n",balance,reflect.TypeOf(balance))
	fmt.Printf("The value is %v and type is %v.\n",boolean,reflect.TypeOf(boolean))
}