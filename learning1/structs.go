package main

import "fmt"

type User struct {
	name   string
	email  string
	status bool
	age    int
}

func main() {
	fmt.Println("Welcome to learn Struct")
	user := User{"Tarek", "tarek@go.dev", true, 21}
	fmt.Println(user)
	fmt.Printf("User details are: %+v\n", user)
	fmt.Printf("Name is %v and Email is %v", user.name, user.email)
}
