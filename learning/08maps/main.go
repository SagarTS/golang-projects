package main

import "fmt"

func main(){

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["PY"] = "Python"


	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	delete(languages,"JS")
	fmt.Println("List of all languages: ", languages)

	//Loops in golang

	for key,value := range languages{
		fmt.Printf("For key %v, value is %v\n",key,value)
	}

}