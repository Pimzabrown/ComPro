package main

import "fmt"

func main() {
	var a int
ReadInput:
	fmt.Print("TYPE NUMBER :")
	fmt.Scan(&a)
	if a < 150 {
		goto ReadInput
	}
	fmt.Println(a)
}