package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://fieldnation.com",
		"http://github.com",
		"http://linkedin.com",
		"http://amazon.com",
	}

	myChannel := make(chan string)

	for _, link := range links {
		go checkLink(link, myChannel)
	}

	// for link := range myChannel {
	// 	go checkLink(link, myChannel)
	// }

	for l := range myChannel {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, myChannel)
		}(l)
	}

}

func checkLink(link string, myChannel chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down!")
		myChannel <- link
		return
	}
	fmt.Println(link, " is up!")
	myChannel <- link
}
