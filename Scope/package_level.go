//Package level Scope
package main

import "fmt"

var x int = 39

func main() {
	fmt.Println(x)
	foo()
}
func foo() {
	fmt.Println(x)
}
