package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Food:")

	// comma ok || err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)

	//conversion from string to float64

	numRating,_ := strconv.ParseFloat(strings.TrimSpace(input),64)
	fmt.Println("Number rating", numRating+1)
	fmt.Printf("Type of  this rating: %T\n", numRating)
}