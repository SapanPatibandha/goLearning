package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "This is test content for file"

	filePath := "../myFile.txt"
	file, err := os.Create(filePath)

	checkError(err)

	length, err := io.WriteString(file, content)

	checkError(err)

	fmt.Println("Length is :", length)

	defer file.Close()

	readFile(filePath)
}

func readFile(filepath string) {

	dataByte, err := ioutil.ReadFile(filepath)

	checkError(err)

	fmt.Println("Text data in file is: \n", string(dataByte))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
