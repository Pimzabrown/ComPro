package main

import "fmt"

func main() {

	add := func(x, y int) int {
		return x + y
	}
	x := add(789, 4122)
	fmt.Println(x)
}