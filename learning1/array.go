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
	nameList = append(nameList, "Nipu", "other")
	fmt.Println("name list is", len(nameList))

	for key, value := range nameList {
		fmt.Println(key, value)
	}

	for i := 0; i < len(nameList); i++ {
		fmt.Println(nameList[i])
	}
}
