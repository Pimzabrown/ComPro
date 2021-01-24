package main

import "fmt"

func main(){
	A := make(map[string]int)
	fmt.Println(len(A))
	A["key"] = 99
	fmt.Println(len(A))
	fmt.Println(A)

	B := map[string]int{"key1": 11, "key2": 1999, "key3": 22}
	B["key4"] = 22111999
	fmt.Println(B)
}