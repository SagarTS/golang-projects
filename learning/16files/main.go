package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main(){
	fmt.Println("Files in golang")
	content := "This will go in the file"

	file,err := os.Create("./learning.txt")

	// if err != nil {
	// 	panic (err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file,content)
	checkNilErr(err)

	fmt.Println("Length is ", length)
	defer file.Close()
	readFile("./learning.txt")
}

func readFile(filename string){
	databyte, err := ioutil.ReadFile(filename)

	checkNilErr(err)

	fmt.Println("Text data inside the file is: \n", string(databyte))
}

func checkNilErr(err error){
	if err != nil {
		panic(err)
	}
}