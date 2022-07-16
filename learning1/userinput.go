package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcomeTitle := "Welcome to user input"
	fmt.Println(welcomeTitle)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\tPut your rating for the order: \r \r abc")

	// comma ok
	input, err := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input, err)
	// input, prefix, err := reader.ReadLine()
	// fmt.Println("Thanks for rating, ", input, prefix, err)
	fmt.Printf("type of this rating is %T \n", input)

	ratingNumber, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err == nil {
		fmt.Println("Adding 1 to your rating: ", ratingNumber+1)
	} else {
		fmt.Println(err)
	}

}
