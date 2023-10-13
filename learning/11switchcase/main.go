package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spot ahead")
	case 3:
		fmt.Println("You can move 3 spot ahead")
		fallthrough
	case 4:
		fmt.Println("You can move 4 spot ahead")
	case 5:
		fmt.Println("You can move 5 spot ahead")
	case 6:
		fmt.Println("You can move 6 spot ahead and roll dice again")
	default:
		fmt.Println("Nothing! Just Default")
	}

}