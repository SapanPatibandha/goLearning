package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	//coma ok || coma error syntex.

	inptu, _ := reader.ReadString('\n')
	fmt.Println("This detail is: ", inptu)

}
