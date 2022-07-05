package main

import (
	"fmt"
	"log"
	"net/http"
	// "html"
)

func main() {
	fmt.Println("Init Main")
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Request URL: %s", req.URL)
		fmt.Fprintf(res, "Request Path: %s", req.URL.Path)
	})
	http.HandleFunc("/test", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "test endpoint.....")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
