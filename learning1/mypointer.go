package main

import "fmt"

func main() {
	fmt.Println("Welcome to the pointer")
	var myptr *int
	fmt.Println("Myptr value is", myptr)

	myNumber := 5
	var ptr = &myNumber
	fmt.Println("Value of the actual pointer is", ptr)
	fmt.Println("Value of the actual pointer is", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is", *ptr)
}
