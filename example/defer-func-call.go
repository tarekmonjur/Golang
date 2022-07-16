package main

import "fmt"

func calculateTaxes(revenue, deductions, credits float64) float64 {
	defer fmt.Println("Taxes Calculated!")
	taxRate := .06143
	fmt.Println("Calculating Taxes")

	if deductions == 0 || credits == 0 {
		return revenue * taxRate
	}

	taxValue := (revenue - (deductions * credits)) * taxRate
	if taxValue >= 0 {
		return taxValue
	} else {
		return 0
	}
}

func main() {
	calculateTaxes(100, 12, 10)
}
