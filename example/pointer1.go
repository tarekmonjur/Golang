package main

import "fmt"

func addHuandred(num *int) {
	*num += 100
}

func brainwash(saying *string) {
	*saying = "Good morning"
}

func main() {
	x := 1
	addHuandred(&x)
	fmt.Println(x)

	greeting := "Hello world"
	brainwash(&greeting)
	fmt.Println(greeting)
}
