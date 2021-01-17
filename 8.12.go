package main

import "fmt"

func main() {
	alphabets := [9]string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	fmt.Println(alphabets)

	A := alphabets[1:8]
	fmt.Println(A)

	B := A[4:5]
	fmt.Println(B)

	C := B[0:7]
	fmt.Println(C)

	Z[0] = "A"
	fmt.Println(alphabets)
}