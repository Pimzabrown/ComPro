package main

import "fmt"

func main() {
	letters := [17]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "P", "Q", "R"}
	x := letters[:]
	y := letters[:9]
	z := letters[6:]
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}