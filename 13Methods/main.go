package main

import "fmt"

func main() {
	fmt.Println("Welcome to methods in golang")

	sapan := User{1, "Sapan", "Patibandha", 45, "abc@abc.com"}

	sapan.GetStatus()

	fmt.Println("New email is :", sapan.GetEmail())
}

type User struct {
	id        int
	firstName string
	lastName  string
	age       int
	email     string
}

func (u User) GetStatus() {
	fmt.Println("Name of user is: ", u.firstName)
}

func (u User) GetEmail() string {
	return u.email
}

// type opt struct {
// 	shortnm      byte
// 	longnm, help string
// 	needArg      bool
// }

// var basenameOpts []opt

// func init() {
// 	basenameOpts = []opt{
// 		opt{
// 			shortnm: 'a',
// 			longnm:  "multiple",
// 			needArg: false,
// 			help:    "Usage for a",
// 		},
// 		opt{
// 			shortnm: 'b',
// 			longnm:  "b-option",
// 			needArg: false,
// 			help:    "Usage for b",
// 		},
// 	}
// }
