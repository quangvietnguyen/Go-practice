package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://www.google.com",
		"https://www.amazon.com",
		"https://vnexpress.net",
		"https://dantri.com.vn",
		"https://kenh14.vn",
		"https://www.facebook.com",
		"https://www.apple.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link,c)
	}

	for l := range c {
		
		go func(link string) {
			time.Sleep(2 * time.Second)
			checkLink(link,c)
		}(l)
	}

	// for  {
	// 	go checkLink(<-c,c)
	// }
	// fmt.Println(<- c)
	// fmt.Println(<- c)
	// fmt.Println(<- c)
	// fmt.Println(<- c)
	// fmt.Println(<- c)
	// fmt.Println(<- c)
	// fmt.Println(<- c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}