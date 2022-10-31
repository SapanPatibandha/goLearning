package main

import (
	"encoding/json"
	"fmt"
)

type cource struct {
	Name     string   `json:"courcename"`
	Prize    int      `json:"prize"`
	Platform string   `json:"platform"`
	Password string   `json:"-"`
	Tages    []string `josn:"tages,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON trial")

	EncodeJSON()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func EncodeJSON() {
	listOfCources := []cource{
		{"C", 100, "Offlie", "password01", []string{"basic", "learning"}},
		{"C++", 200, "Offlie", "password02", []string{"test", "learning"}},
		{"C#", 300, "Online", "password03", nil},
	}

	finalJSON, err := json.MarshalIndent(listOfCources, "", "\t")
	// finalJSON, err := json.Marshal(listOfCources)

	checkError(err)

	fmt.Printf("%s \n", finalJSON)
}
