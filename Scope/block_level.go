//Block Level /Function Scope
package main

import "fmt"

var y int = 539

func main() {
	x := 39
	fmt.Println(y)
	fmt.Println(x)
	foo()
}
func foo() {
	fmt.Println(y)
	//fmt.Println(x) Wont Execute this cos its block level
}
