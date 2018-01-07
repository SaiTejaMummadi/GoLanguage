//Closure
package main

import "fmt"

func main() {
	x := 39
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "This y can be accessed only in the block"
		fmt.Println(y)
	}
	//fmt.Println(y) y wont work out of the scope
}
