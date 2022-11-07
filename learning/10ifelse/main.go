package main

import "fmt"

func main(){

	loginCount := 20
	var result string

	if loginCount < 10 {
		result = "Resular user"
	} else if loginCount < 20 {
		result = "Watch out"
	} else {
		result = "Else case"
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is greater than 10")
	}
}

