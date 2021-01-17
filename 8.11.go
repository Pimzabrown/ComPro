package main

import "fmt"

func main() {
	A := make([]int, 5)
	fmt.Println(A)
	fmt.Println(len(A))
	fmt.Println(cap(A))

	B := make([]int, 40, 42)
	fmt.Println(B)
	fmt.Println(len(B))
	fmt.Println(cap(B))
}