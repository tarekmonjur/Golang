package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	ContactInfo
}

type ContactInfo struct {
	email   string
	zipCode int
}

func main() {
	user := Person{
		firstName: "Tarek",
		age:       30,
		ContactInfo: ContactInfo{
			email:   "tarek@example.com",
			zipCode: 12345,
		},
	}

	if user.lastName == "" {
		user.lastName = "Ahammed"
	}

	// monjur := &user
	// monjur.setFirstName("Monjur")
	// monjur.print()

	user.setFirstName("Monjur")
	user.print()
}

func (u Person) print() {
	fmt.Printf("%+v\n", u)
}

func (u *Person) setFirstName(firstName string) {
	u.firstName = firstName
}
