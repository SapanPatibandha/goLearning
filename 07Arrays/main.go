package main

import "fmt"

func main() {
	fmt.Println("Welcome to the class of Arrays. ")

	var frutlist [4]string

	frutlist[0] = "apple"
	frutlist[1] = "Orange"
	frutlist[3] = "banana"

	fmt.Println("Frutlist is :", frutlist)
	fmt.Println("Frutlist is :", len(frutlist))

}
