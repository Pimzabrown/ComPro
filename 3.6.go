package main

import "fmt"

func main() {
	name :=[]string{}
	name = append(name, "MARKLEE IS SO CUTE ")
	letters := [7]string{"M", "A", "R", "K", "L", "E", "E"}
	A := letters[:]
	B := letters[:4]
	C := letters[6:]
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(name)
}

