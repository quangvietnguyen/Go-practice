package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct{
	height float64
	base float64
}

type square struct{
	sideLength float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape){
	fmt.Println(s.getArea())
}

func main() {
	s := square{3}
	t := triangle{2,3}
	printArea(s)
	printArea(t)
}