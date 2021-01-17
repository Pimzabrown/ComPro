package main

import "fmt"

func main() {
	fmt.Println(1356 * 5236)
	fmt.Println((-5 * 10.8) / 9 + (10 % 2))
	fmt.Println((true && false) || (!true))
	fmt.Printf("10 is of type %T\n", 10)
	fmt.Printf("10 is of type %T\n", string(10))
	fmt.Printf("10 is of type %T\n", float64(10))
}