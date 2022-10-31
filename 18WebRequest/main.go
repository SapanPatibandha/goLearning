package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web Request --")

	// getRequestTrial("http://localhost:8000/get")

	// fmt.Println("==========================================================================")

	//postRequestTrial("http://localhost:8000/post")

	// fmt.Println("==========================================================================")

	performFormRequest("http://localhost:8000/postform")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getRequestTrial(url string) {

	response, err := http.Get(url)

	checkError(err)

	defer response.Body.Close()

	fmt.Println("Status code :", response.Status)
	fmt.Println("Content length is :", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println("Content is :", string(content))

	var responseString strings.Builder

	byteCount, _ := responseString.Write(content)

	fmt.Println("Bite Count is :", byteCount)

	fmt.Println("Data is : ", responseString.String())

}

func postRequestTrial(url string) {

	//create json format
	requestBody := strings.NewReader(`
	{
		"FirstName" : "Sapan", 
		"LastName" : "Patibandha", 
		"Age" : 45
	}
	`)

	response, err := http.Post(url, "application/json", requestBody)

	checkError(err)

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}

func performFormRequest(myurl string) {

	data := url.Values{}
	data.Add("firstname", "Sapan")
	data.Add("lasttname", "Sapan")
	data.Add("email", "patibandha@yahoo.com")

	response, err := http.PostForm(myurl, data)

	checkError(err)

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
