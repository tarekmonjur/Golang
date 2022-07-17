package main

import "fmt"

func main() {
	fmt.Println("Welcome to loop,  continue, break, goto")
	days := []string{"Friday", "Saturday", "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday"}
	fmt.Println(days)

	fmt.Println("Loop one..........")
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	fmt.Println("Loop two..........")
	for j := range days {
		fmt.Println(days[j])
	}

	fmt.Println("Loop three.........")
	for index, day := range days {
		fmt.Println(index, day)
	}

	i := 0
	for i < len(days) {
		if i == 6 {
			break
			goto jump1
		}

		if i == 3 {
			i++
			continue
		}
		fmt.Println(i)
		i++
	}

jump1:
	fmt.Println("jum here from loop")
}
