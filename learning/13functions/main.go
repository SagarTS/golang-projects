package main

import "fmt"

func main(){
	fmt.Println("Functions")
	greeter()

	result := adder(3,5)

	fmt.Println("Result is: ", result)

	proResult, str := proAdder(1,2,3,4,5,6,7,8,9,10)
	fmt.Println(proResult,str)
}

func adder (valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder (values ...int) (int, string){
	total := 0

	for _,value := range values{
		total += value
	}

	return total, "Hello"
}

func greeter(){
	fmt.Println("Namaste from greeter function")
}