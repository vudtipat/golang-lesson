package main
import "fmt"

func main()  {
	var str1,str2 string = "Hello","World"

	fmt.Println(str1,str2,"!!!")
	fmt.Print(str1,"\n",str2,"\n")
	fmt.Printf("%s %s\n",str1,str2)

	var name string
	fmt.Print("Enter your nickname: ")
	fmt.Scanf("%s",&name)
	fmt.Println("Hi there, ",name)
	fmt.Print("Enter your firatname lastname: ")
	var fname,lname string
	fmt.Scanf("%s %s",&fname,&lname)
	fmt.Println("Hi there, ",fname,lname)
	
	var (
		s string = "Hello"
		i int = 64
	)

	fmt.Printf("%q\n",s)
	fmt.Println(i)

}