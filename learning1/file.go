package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to file in golang")
	content := "create a new file and write this context to the file"
	file, err := os.Create("./mylogs.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("file length is", length)
	defer file.Close()
	readFile("./mylogs.txt")
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	checkNilError(err)
	fmt.Println(string(dataByte))
}
