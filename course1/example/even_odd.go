package main

import (
	"fmt"
	"strings"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenNumbers, oddNumbers := evenOdd(numbers)

	fmt.Println("Even Numbers: ", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(evenNumbers)), ", "), "[]"))
	fmt.Println("Odd Numbers: ", strings.Trim(strings.Replace(fmt.Sprint(oddNumbers), " ", ", ", -1), "[]"))

	// for even := range evenNumbers {
	// 	fmt.Print(even)
	// }

	// for odd := range oddNumbers {
	// 	fmt.Print(odd)
	// }
}

func evenOdd(numbers []int) ([]int, []int) {
	var even []int
	var odd []int

	for _, num := range numbers {
		if num%2 == 0 {
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
	}
	return even, odd
}
