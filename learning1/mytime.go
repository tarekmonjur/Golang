package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to my time of golang")
	presentTime := time.Now()
	fmt.Println("Present Time:", presentTime)
	fmt.Printf("present time type %T \n", presentTime)
	fmt.Println(presentTime.Format("01/02/2006"))

	createDate := time.Date(presentTime.Year(), time.September, presentTime.Day(), 23, 23, 0, 0, time.UTC)
	fmt.Println(createDate.Format("01/02/2006 Monday 15:04:05"))
}
