package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["peter"] = 32
	ages["anne"] = 35

	// fmt.Println(ages["viet"])

	// ages["viet"] += 1
	// fmt.Println(ages["viet"])

	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}