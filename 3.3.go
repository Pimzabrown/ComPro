package main

import "fmt"

func main() {
	A := make([]int, 10)
	fmt.Println(A)
	fmt.Println(len(A))
	fmt.Println(cap(A))

	R := make([]int, 10,20)
	fmt.Println(R)
	fmt.Println(len(R))
	fmt.Println(cap(R))
}