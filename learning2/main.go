package main

import "fmt"

func main() {
	// fmt.Println("Hello, World!")

	// var a int
	// fmt.Println(a)

	// var b string
	// fmt.Println(b)

	// var c bool
	// fmt.Println(c)

	// var p *int
	// fmt.Println(p)
	// p = &a
	// fmt.Println(p)
	// fmt.Println(*p)

	// // fmt.Printf("Tarek %s", "Ahammed")
	// name := "Tarek"
	// // name = fmt.Sprintf("%s Ahammed", name)
	// // fmt.Println(name)

	// var age int
	// fmt.Print("Enter your name: ")
	// // fmt.Scanf("%s", &name)
	// fmt.Scan(&name)
	// fmt.Print("Enter your age: ")
	// fmt.Scan(&age)
	// fmt.Println("Your name is", name, "and age is", age)
	// message := fmt.Sprintf("Your name is %s, age is %d \n", name, age)

	// reader, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// fmt.Println(reader)

	// fields := strings.Fields(reader)
	// fmt.Println(fields)

	// myName := strings.Join(fields[:len(fields)-1], " ")
	// fmt.Println(myName)

	// myAge := fields[len(fields)-1]
	// fmt.Println(myAge)

	// day2
	// file, _ := os.Create("test.txt")
	// fmt.Fprintln(file, message)
	// file.Close()

	// ara := [5]int{1, 2, 3, 4, 5}
	// fmt.Println(ara)
	// slice := ara[:3]
	// slice[0] = 100
	// fmt.Println(slice, len(slice), cap(slice))
	// fmt.Println(ara, len(ara), cap(ara))

	// slice2 := []int{1, 2, 3}
	// slice2 = append(slice2, 2, 3, 5, 6, 7)
	// fmt.Println(slice2, len(slice2), cap(slice2))

	// slice3 := make([]int, 3, 5)
	// print(slices.Max(slice3))
	// fmt.Println(slice3, len(slice3), cap(slice3))

	// slice4 := []int{1, 2, 3, 4, 5}
	// slice4 = append(slice4[:2], slice4[3:]...)
	// fmt.Println(slice4, len(slice4), cap(slice4))

	map1 := make(map[string]int)
	fmt.Println(map1)

	map2 := map[string]int{
		"one": 1,
		"two": 2,
	}
	map2["test"] = 111
	fmt.Println(map2, map2["three"])
	map2["three"] = 3
	mapValue, ok := map2["three"]
	fmt.Println(map2, mapValue, ok, map2["three"])
}
