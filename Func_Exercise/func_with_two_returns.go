package main

import "fmt"

func divi(n int) (int, bool) {
	return n / 2, n%2 == 0
}
func main() {
	x, y := divi(4)
	fmt.Println(x, y)
}
