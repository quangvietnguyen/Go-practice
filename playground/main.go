package main

import (
	"fmt"
)

func main() {
	const name, age = "Kim", 22
	n,err := fmt.Println(name, "is", age, "years old.")
	if err != nil {
		fmt.Println(err)
	}
fmt.Println("Number of bytes written:", n)
	// It is conventional not to worry about any
	// error returned by Println.

}
