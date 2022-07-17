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
	fmt.Printf("Name is %v and Email is %v\n", user.name, user.email)
	user.getStatus()
	user.setNewEmail("monjur@go.dev")
	fmt.Printf("Name is %v and Email is %v\n", user.name, user.email)
}

func (u User) getStatus() {
	fmt.Println("Is user active:", u.status)
}

func (u User) setNewEmail(email string) {
	u.email = email
	fmt.Println("User email is:", u.email)
}
