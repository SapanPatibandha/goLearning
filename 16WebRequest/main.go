package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev" //"https://www.google.com/" //

func main() {
	fmt.Println("Web request")

	response, err := http.Get(url)

	checkError(err)

	fmt.Printf("Response from web site %T\n", response)

	defer response.Body.Close()

	dataBytes, err := io.ReadAll(response.Body)

	checkError(err)

	fmt.Println("data on site is : ", string(dataBytes))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
