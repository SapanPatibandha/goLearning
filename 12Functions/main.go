package main

import "fmt"

func main() {
	fmt.Println("Welcome to function in goLang")
	greater()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	a, b, c := adderTwo(2, 3, 4, 5, 9)

	fmt.Println("Response from dynameic param values: ", a)
	fmt.Println("Response from dynameic param values: ", b)
	fmt.Println("Response from dynameic param values: ", c)
}

func greater() {
	fmt.Println("Hellow from function of golang")
}

func adder(firstVal int, secondVal int) int {
	return firstVal + secondVal
}

func adderTwo(vlaues ...int) (int, string, string) {
	total := 0

	for _, val := range vlaues {
		total += val
	}

	return total, "text message", "abcde"
}
