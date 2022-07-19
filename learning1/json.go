package main

import (
	"encoding/json"
	"fmt"
)

type posts struct {
	Title  string   `json:"title"`
	Body   string   `json:"body"`
	UserId int      `json:"-"`
	Tags   []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to learn json in Golang")
	postData := []posts{
		{"title one", "body...sadfaf", 1, []string{"tag1", "tag2"}},
		{"title two", "body...here", 2, nil},
	}
	encodeDataB := EncodeJson(postData)
	// decodeData := DecodeJson(encodeDataB).([]posts)
	decodeData := DecodeJson(encodeDataB)
	fmt.Println(decodeData[0])

}

func EncodeJson(data interface{}) []byte {
	jsonData, err := json.Marshal(data)
	// jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(jsonData))
	// fmt.Printf("%s\n", jsonData)
	return jsonData
}

func DecodeJson(data []byte) []interface{} {
	var newJson []interface{}
	// var newJson []any
	// var newJson []posts
	isValid := json.Valid(data)
	if isValid {
		json.Unmarshal(data, &newJson)
		// fmt.Println(newJson)
		// fmt.Printf("%#v\n", newJson)
		// fmt.Println(newJson[0].Title)
	} else {
		fmt.Println("Json was not valid")
	}
	return newJson
}
