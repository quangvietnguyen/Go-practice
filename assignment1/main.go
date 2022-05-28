package main

import "fmt"

func main() {
	num := []int{}

	for i := 0; i <= 10; i++ {
		num = append(num,i)
	}

	for _, n := range num {
		if n % 2 == 0 {
			fmt.Println(n,"is even")
		} else {fmt.Println(n,"is odd")}
	}
}