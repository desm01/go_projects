package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
		"http://stackoverflow.com",
	}

	c := make(chan string)

	for _, link := range urls {
		go getStatus(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			getStatus(link, c)
		}(l)
	}
}

func getStatus(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
		c <- link
		return
	}
	fmt.Println(link, "is currently ok")
	c <- link
}
