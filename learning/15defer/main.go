package main

import "fmt"

func main(){
	defer fmt.Println("Last line")
	defer fmt.Println("World")
	fmt.Println("Hello")

	myDefer()
}

func myDefer(){
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}