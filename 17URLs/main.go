package main

import (
	"fmt"
	"net/url"
)

const URL = "https://lco.dev/test?test=abc&test2=343"

func main() {
	fmt.Println("welcome to URL handeling")

	fmt.Println(URL)

	//parsing URL
	result, err := url.Parse(URL)

	checkError(err)

	fmt.Println("Result :", result.Scheme)
	fmt.Println("Result :", result.Host)
	fmt.Println("Result :", result.Hostname())

	fmt.Println("Result :", result.Path)

	fmt.Println("Result :", result.User.Username())

	qParam := result.Query()

	fmt.Println("Query param :", qParam)

	for _, val := range qParam {
		fmt.Println("Param is : ", val)
	}

	partsOfURL := &url.URL{
		Scheme:  "https",
		Host:    "google.com",
		Path:    "/map",
		RawPath: "UserName=sapna",
	}

	fmt.Println("Created url is: ", partsOfURL.String())
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
