package main

import "fmt"

func main() {
	fmt.Println("Welcome to learn methods")
	Greetings()
	Greeter()
	var total, message = adder(3, 5, 7, 9)
	fmt.Println(message, total)
}

func Greetings() {
	fmt.Println("Hello go method")
}

func Greeter() {
	fmt.Println("Another metnod")
}

func adder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "final result is"
}
