package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to example of slices")

	var frutList = []string{"Apple", "banana", "Orange"}

	fmt.Printf("Type of frut list is %T \n", frutList)

	fmt.Println(frutList)
	frutList = append(frutList, "test1", "test 2")
	fmt.Println(frutList)

	frutList = append(frutList[1:4])

	fmt.Println(frutList)

	fmt.Println("Lenght of slice: ", len(frutList))

	//----------------------------------------------------------------------

	scores := make([]int, 4)

	scores[0] = 2
	scores[1] = 9
	scores[2] = 7
	scores[3] = 1

	fmt.Println(scores)

	sort.Ints(scores)

	fmt.Println(scores)

}
