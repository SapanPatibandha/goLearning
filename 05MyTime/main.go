package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study for golang.")

	myTime := time.Now()

	fmt.Println(myTime)

	fmt.Println(myTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2020, 1, 18, 0, 0, 0, 0, time.Now().UTC().Location())
	fmt.Println("Create Date: ", createDate)
}
