package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to the http web request")
	makeGetRequest("https://jsonplaceholder.typicode.com/posts/1")
	// makePostRequest("https://jsonplaceholder.typicode.com/posts")
}

func makeGetRequest(endpoint string) {
	var res, err = http.Get(endpoint)
	if err != nil {
		panic(err)
	}

	fmt.Printf("response type is %T\n", res)
	fmt.Println(res.Status, res.StatusCode, res.Close)
	defer res.Body.Close()
	dataByte, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	// result := string(dataByte)
	// fmt.Println("result: ", result)
	// fmt.Println(res.Close)

	var resString strings.Builder
	resString.Write(dataByte)
	fmt.Println(resString.String())
}

func makePostRequest(url string) {
	payload := strings.NewReader(`
		{
			"title": "foo",
			"body": "bar",
			"userId": 1
		}
	`)

	res, err := http.Post(url, "application/json", payload)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	content, _ := ioutil.ReadAll(res.Body)
	data := string(content)
	fmt.Println("data", data)

}

func makePostFormRequest(endpoint string) {
	payload := url.Values{}
	payload.Add("title", "foo")
	payload.Add("body", "bar")

	res, err := http.PostForm(endpoint, payload)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	content, _ := ioutil.ReadAll(res.Body)
	data := string(content)
	fmt.Println("data", data)

}
