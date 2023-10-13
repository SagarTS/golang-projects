package main

import "fmt"

func main(){
	// no inheritance in golang

	john := User {
		"John", "john@go.com", true, 40,
	}

	fmt.Println(john)
	fmt.Printf("John's details are: %+v\n",john)
	fmt.Printf("Name is %v, and email is %v\n",john.Name, john.Email)
}

type User struct {
	Name	string
	Email	string
	Status	bool
	Age		int
}