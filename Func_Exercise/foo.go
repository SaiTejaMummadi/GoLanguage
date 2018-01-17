package main

import "fmt"

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4, 5}
	foo(aSlice...)
	foo()
}
func foo(n ...int) {
	fmt.Println(n)
}
