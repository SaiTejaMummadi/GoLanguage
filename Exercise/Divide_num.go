package main

import "fmt"

func main() {
	var numOne int
	var numTwo int
	fmt.Print("Enter big Number: ")
	fmt.Scan(&numOne)
	fmt.Print("Enter small Number: ")
	fmt.Scan(&numTwo)
	fmt.Println(numOne, "/", numTwo, "=", numOne/numTwo)
}
