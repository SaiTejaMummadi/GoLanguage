//Variable Shadowing
//Bad Practice Use unique names for Variables and functions

package main

import "fmt"

func max(x int) int {
	return 42 + x
}
func main() {
	max := max(7)
	fmt.Println(max) //Max is now the variable.
}
