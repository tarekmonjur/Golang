package main

import (
	"encoding/json"
	"fmt"
)

type posts struct {
	Title  string   `json:"title"`
	Body   string   `json:"body"`
	UserId int      `json:"-"`
	Tags   []string `json:"tags","omitempty"`
}

func main() {
	fmt.Println("Welcome to learn json in Golang")
	postData := []posts{
		{"title one", "body...sadfaf", 1, []string{"tag1", "tag2"}},
		{"title two", "body...here", 2, nil},
	}
	// jsonData, err := json.Marshal(postData)
	jsonData, err := json.MarshalIndent(postData, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}
