package main

import "fmt"

func main() {
	fmt.Println("Welcome to Differ functionality 1")

	defer fmt.Println("This is first")
	fmt.Println("Welcome to Differ functionality 2")
	defer fmt.Println("this is second line")

	fmt.Println("Welcome to Differ functionality 3")
	defer fmt.Println("this is third line")

	myDifer()
}

func myDifer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
