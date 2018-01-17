package main

import "fmt"

func max(num ...int) int {
	fmt.Printf("%T\n", num)
	var large int
	for _, v := range num {
		if v > large {
			large = v
		}
	}
	return large
}
func main() {
	great := max(4, 7, 9, 3, 5, 8, 6, 7, 23, 45, 3, 4, 6, 7)
	fmt.Println(great)
}
