package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Golang array")

	var fruitList [5]string
	fruitList[0] = "Mango"
	fruitList[1] = "Orange"
	fruitList[2] = "Banana"
	fmt.Println("My fruit list is", fruitList)
	fmt.Println("My fruit count is", len(fruitList))

	var nameList = []string{"Muntasir", "Muntaha", "Tarek", "Monjur"}
	fmt.Println("name list is", len(nameList))
}
