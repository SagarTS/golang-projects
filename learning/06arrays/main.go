package main

import "fmt"

func main(){

	var superHeros [4]string

	superHeros[0] = "SuperMan"
	superHeros[1] = "Batman"
	superHeros[3] = "Spiderman"

	fmt.Println("SuperHeros are ",superHeros)
	fmt.Println("SuperHeros are ",len(superHeros))

	var superVillians = [3]string{"BlackAdam","Dr.Doom","Galactus"}
	
	fmt.Println("Super Villians are ", superVillians)

}