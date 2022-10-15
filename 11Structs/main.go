package main

import "fmt"

func main() {
	fmt.Println("Welcome to get Structure. ")

	sapan := User{1, "Sapan", "Patibandha", 45}

	fmt.Println("User detail is: ", sapan)

	fmt.Printf("User detail is: %+v \n", sapan)

	fmt.Printf("User Name is %v %v and age is %v \n", sapan.firstName, sapan.lastName, sapan.age)
}

type User struct {
	id        int
	firstName string
	lastName  string
	age       int
}
