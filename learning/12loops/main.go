package main

import "fmt"

func main(){
	days := []string{"Sunday","Monday","Tuesday","Friday"}

	fmt.Println(days)

	for i:=0; i < len(days); i++{
		fmt.Println(days[i])
	} 

	for i := range days{
		fmt.Println(days[i])
	}

	for _, day := range days{
		fmt.Printf("value is %v\n",day)
	}

	rougueValue := 1

	for rougueValue < 10 {
		if(rougueValue == 2){
			goto here
		}

		if rougueValue == 5{
			rougueValue++
			continue
		}
		fmt.Println("Value is: ",rougueValue)
		rougueValue++
	}

	here: 
		fmt.Println("You jumped here.")

}	
