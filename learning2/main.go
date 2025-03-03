package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")

	var a int
	fmt.Println(a)

	var b string
	fmt.Println(b)

	var c bool
	fmt.Println(c)

	var p *int
	fmt.Println(p)
	p = &a
	fmt.Println(p)
	fmt.Println(*p)

	// fmt.Printf("Tarek %s", "Ahammed")
	name := "Tarek"
	// name = fmt.Sprintf("%s Ahammed", name)
	// fmt.Println(name)

	var age int
	fmt.Print("Enter your name: ")
	// fmt.Scanf("%s", &name)
	fmt.Scan(&name)
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)
	fmt.Println("Your name is", name, "and age is", age)
	fmt.Printf("Your name is %s, age is %d \n", name, age)

	reader, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println(reader)

	fields := strings.Fields(reader)
	fmt.Println(fields)

	myName := strings.Join(fields[:len(fields)-1], " ")
	fmt.Println(myName)

	myAge := fields[len(fields)-1]
	fmt.Println(myAge)
}
