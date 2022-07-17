package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to the mySlices")
	var fruitList = []string{}
	fmt.Println("My fruit list init", len(fruitList))
	fmt.Printf("Type of fruit list is %T\n", fruitList)

	fruitList = append(fruitList, "Apple", "Mango", "Banana", "Orange")
	fmt.Println("my fruit list is", fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println("my fruit list now is", fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println("my new fruit list is", fruitList)

	var scores = make([]int, 5)
	scores[0] = 222
	scores[1] = 555
	scores[2] = 333
	scores[3] = 444
	scores[4] = 777
	// scores[5] = 999
	scores = append(scores, 999, 1000)
	sort.Ints(scores)
	fmt.Println("Score value is", scores)
	fmt.Println(sort.IntsAreSorted(scores))

	for i := range scores {
		fmt.Println(scores[i])
	}

	// remove item from slices
	var courses = []string{"php", "javascript", "python", "java", "go"}
	var index = 3
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
