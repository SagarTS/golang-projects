package main

import (
	"fmt"
	"sort"
)

func main() {
	var superHeroes = []string{"SuperMan","Batman","Spiderman"}

	superHeroes = append(superHeroes,"Ironman","Hulk")

	fmt.Printf("Type of superHeroes is %T\n", superHeroes)
	fmt.Println("SuperHeroes are ", superHeroes)

	superHeroes = append(superHeroes[1:])
	fmt.Println("SuperHeroes are ", superHeroes)

	fmt.Println(sort.StringsAreSorted((superHeroes)))
	sort.Strings(superHeroes)
	fmt.Println(superHeroes)

	var index int = 2
	superHeroes = append(superHeroes[:index],superHeroes[index+1:]...)
	fmt.Println(superHeroes)
}