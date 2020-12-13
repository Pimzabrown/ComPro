package main

import "fmt"

func main() {
	A := make(map[string]int)
	fmt.Println(len(A))
	A["key"] = 127
	fmt.Println(A)
	fmt.Println(len(A))
	fmt.Println(A)

	B := map[string]int{"key1": 50, "key2": 70}
	B["key3"] = 90
	fmt.Println(B)
}