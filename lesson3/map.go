package main

import (
	"fmt"
)

func main() {
	mmap := map[string]string{"en":"English","th":"Thai"}
	fmt.Println(mmap)
	fmt.Println(len(mmap))

	value,isFound := mmap["jp"]
	fmt.Println(isFound)
	if(isFound) {
		fmt.Println(value)
	}

	mmap["jp"] = "Japan"
	fmt.Println(mmap)

	delete(mmap,"jp")
	fmt.Println(mmap)

	for key, value := range mmap {
		fmt.Println(key," => ",value)
	}
}