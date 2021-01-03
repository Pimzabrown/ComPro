package main

import "fmt"

func main() {

	add := func(x, y int) int {
		return x * y + 2
	}
	x := add(15, 40)
	fmt.Println(x)
}