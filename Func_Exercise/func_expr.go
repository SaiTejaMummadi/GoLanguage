package main

import "fmt"

func main() {
	divi := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}
	x, y := divi(4)
	fmt.Println(x, y)
}
