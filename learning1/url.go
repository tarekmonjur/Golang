package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://tarekmonjur.com:8000/users?name=tarek&age=30"

func main() {
	fmt.Println("Welcome to url parsing")
	fmt.Println(myUrl)

	result, _ := url.Parse(myUrl)
	fmt.Println("Host:", result.Host)
	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Path:", result.Path)
	fmt.Println("Port:", result.Port())
	fmt.Println("RawQuery:", result.RawQuery)
	fmt.Println("Query:", result.Query())
	qparams := result.Query()
	fmt.Println("Name:", qparams["name"])
	fmt.Println("Age:", qparams.Get("age"))

	makeUrl := &url.URL{
		Scheme:   "https",
		Host:     "tarekmonjur.com",
		Path:     "todos",
		RawQuery: "id=111",
	}
	var newUrl string = makeUrl.String()
	fmt.Println("new url is:", newUrl)

}
