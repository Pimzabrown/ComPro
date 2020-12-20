package main

import "fmt"

func main() {
	var A int
ReadInput:
	fmt.Print("Type Number:")
	fmt.Scan(&A)
	if A < 50 {
		goto ReadInput
	}
	fmt.Println(A)
}