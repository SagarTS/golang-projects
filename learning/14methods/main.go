package main

import "fmt"

func main(){
	john := User {
		"John", "john@go.com", true, 40,
	}

	fmt.Println(john)
	fmt.Printf("John's details are: %+v\n",john)
	fmt.Printf("Name is %v, and email is %v\n",john.Name, john.Email)

	john.GetStatus()
	john.NewMail()
}

type User struct {
	Name	string
	Email	string
	Status	bool
	Age		int
}

func (u User) GetStatus(){
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail(){
	u.Email = "new@mail.com"
	fmt.Println("Email of this user is: ", u.Email)
}