package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	fmt.Println(&file)
	if (err != nil) {
		fmt.Println(err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}