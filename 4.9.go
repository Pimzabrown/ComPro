package main

import "fmt"

func main() {
	score := 75
	if score > 80 {
		fmt.Println("A")
	}
	if score > 70 {
		fmt.Println("B")
	}
	if score > 60 {
		fmt.Println("C")
	}
	if score > 50 {
		fmt.Println("D")
	}
	if score > 0 {
		fmt.Println("F")
	}
}