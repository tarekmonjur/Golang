package main

import "fmt"

const loginToken = "kfja8798qrljq0g98a903rqjl"

func main() {
	var name string = "Tarek Monjur"
	fmt.Printf("%s : variable is of type: %T \n", name, name)

	var isLoggedIn bool = false
	fmt.Printf("%b : variable is of type: %T \n", isLoggedIn, isLoggedIn)

	var price uint16 = 1200
	fmt.Printf("%d : variable is of type: %T \n", price, price)

	var amount float32 = 120.41243141
	fmt.Printf("%f : variable is of type: %T \n", amount, amount)

	// default type
	var defaltvalue int
	fmt.Printf("%d : variable is of type: %T \n", defaltvalue, defaltvalue)

	// implicit type
	var website = "tarekmonjur.com"
	fmt.Printf("%s : variable is of type: %T \n", website, website)

	// no var style or inference type
	numberOfOrder := 500000
	fmt.Printf("%d : variable is of type: %T \n", numberOfOrder, numberOfOrder)

	fmt.Printf("%s : variable is of type: %T \n", loginToken, loginToken)
}
