package main

import "fmt"

func main(){
	fmt.Println("POINTERS")

	number := 4 
	var prt = &number 

	fmt.Println("Value of pointer is ", prt)
	fmt.Println("Value of pointer is ", *prt)

	*prt = *prt + 2
	fmt.Println("New value is ", number)
}