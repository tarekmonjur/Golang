package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const BMI = "BMI = weight(kg) / (height(m) * height(m))"
	var name string
	var weight float32
	var height float32

	fmt.Print("Enter your name: ")
	name, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Print("Enter your weight in kg: ")
	fmt.Scan(&weight)
	fmt.Print("Enter your height in meters: ")
	fmt.Scan(&height)

	fmt.Printf("\n\nWelcome %s! Let's calculate your BMI.\n", name)
	fmt.Println("\nFormula used: ", BMI)

	bmiResult := weight / (height * height)
	fmt.Println("\nYour BMI is: ", bmiResult)

	if bmiResult > 18.5 && bmiResult < 24.9 {
		fmt.Println("Health Status: You are in normal weight range.")
	} else {
		fmt.Println("Health Status: You are not in normal weight range.")

	}
}
