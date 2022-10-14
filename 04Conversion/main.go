package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome to my app.")
	fmt.Println("Please enter values between 1 to 5.")

	reader := bufio.NewReader(os.Stdin)

	inptu, _ := reader.ReadString('\n')

	fmt.Println("Thanks for reading: ", inptu)

	numValue, erro := strconv.ParseFloat(strings.TrimSpace(inptu), 64)

	if erro != nil {
		fmt.Println(erro)
	} else {
		fmt.Println("add one to this.", numValue+1)
	}
}
