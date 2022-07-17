package main

import "fmt"

func main() {
	fmt.Println("Welcome to lear golan map")

	languages := make(map[string]string)
	languages["php"] = "Hypertext Preprocessor"
	languages["py"] = "Python"
	languages["js"] = "Javascript"
	languages["java"] = "Java"
	fmt.Println("Languages:", languages)
	fmt.Println("JS shorts for: ", languages["js"])
	delete(languages, "java")
	fmt.Println("Languages:", languages)

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v \n", key, value)
	}

}
