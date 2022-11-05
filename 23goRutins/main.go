package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to goRutin")

	go greeter("hellow")
	greeter("World")
}

func greeter(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}
