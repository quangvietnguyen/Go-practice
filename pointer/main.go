package main

import "fmt"

func f() *int {
	v:=1
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}

func main() {
	v:=1
	p:=incr(&v)
	fmt.Println(p)
}