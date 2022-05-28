package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hello, world"
	fmt.Println(len(s))     // "12"
	fmt.Println(s[0], s[7]) // "104 119"  ('h' and 'w')

	// c := s[len(s)]
	// fmt.Println(c) // "0"

	fmt.Println(s[0:5]) // "hello"
	fmt.Println(s[0:1])
	fmt.Println("\a")
	fmt.Println(len("Hello, 🌞")) // 11
	fmt.Println(utf8.RuneCountInString("Hello, 🌞")) // 8
	
	for i := 0; i < len("Hello, 🌞"); {
    	r, size := utf8.DecodeRuneInString("Hello, 🌞"[i:])
    	fmt.Printf("%d\t%c\n", i, r)
    	i += size
	}

	for i, r := range "Hello, 🌞" {
    	fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	s = "🌞"
	fmt.Printf("% x\n", s) // f0 9f 8c 9e
	r := []rune(s)
	fmt.Printf("%x\n", r) // [1f31e]
	fmt.Printf("%T\n",s)
	fmt.Println(string(r)) // 🌞

	fmt.Println(string(65))     // "A", not "65"
	fmt.Println(string(0x1f31e)) // "🌞"

	s = "abc"
	b := []byte(s)
	fmt.Printf("% x\n", b) // 61 62 63
	fmt.Println(b)
	s2 := string(b)
	fmt.Println(s2) // abc
}