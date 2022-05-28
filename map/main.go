package main

import "fmt"

func main() {
	// var color map[string]string
	color := map[string]string{
		"red": "#ff0000",
		"green": "#33af333",
		"white": "#ffffff",
	}

	// color := make(map[string]string)

	// color["white"] = "#ffffff"

	// delete(color,"white")

	printMap(color)
}

func printMap(c map[string]string) {
	for color,hex := range c {
		fmt.Println(color, hex)
	}
}