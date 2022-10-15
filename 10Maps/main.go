package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to maps in golang - Dictionary")

	languages := make(map[string]string)

	languages["c#"] = "C Sharp"
	languages["jv"] = "Java"
	languages["py"] = "Python"

	fmt.Println("values of map or dictionary is: ", languages)

	fmt.Println("values for c#: ", languages["c#"])

	delete(languages, "jv")
	fmt.Println("values of map or dictionary is: ", languages)

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v \n", key, value)
	}
}
