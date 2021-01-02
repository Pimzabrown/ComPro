package main

import "fmt"

func substract(x int, y int) int {
	ans := x + y + 10 - 5
	return ans
}

func main() {
	x := substract(87564, 5789)
	fmt.Println(x)
}