package main

import "fmt"

func main() {
	fmt.Println("Welcome to the class of pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 33

	var ptr = &myNumber
	var myNumber2 = ptr

	fmt.Println("Reference to pointer is", ptr)
	fmt.Println("Value of pointer is", *ptr)

	fmt.Println("Value of pointer is", *myNumber2)

	myNumber = myNumber + 2

	fmt.Println("Value of pointer is", *myNumber2)

}
