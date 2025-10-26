package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args
	fileName := args[1]

	// byteContent, err := os.ReadFile(fileName)
	// if err != nil {
	// 	fmt.Println("Error reading file:", err)
	// 	return
	// }
	// fmt.Println(string(byteContent))

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}
