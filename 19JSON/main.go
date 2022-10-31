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

	// EncodeJSON()
	DecodeJSON()
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

func DecodeJSON() {
	jsonData := []byte(`
	{
		"courcename": "C",
		"prize": 100,
		"platform": "Offlie",
		"Tages": ["basic","learning"]
	}
	`)

	var listCource cource

	if json.Valid(jsonData) {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonData, &listCource)
		fmt.Printf("%#v\n", listCource)
	} else {
		fmt.Printf("JSON WAS NOT VALID")
	}

	//convert JSON data to just key value pare.

	var onlineData map[string]interface{}
	json.Unmarshal(jsonData, &onlineData)
	fmt.Printf("%#v\n", onlineData)

	for k, v := range onlineData {
		fmt.Printf("key is %v and values is %v and type is %t\n", k, v, v)
	}
}
