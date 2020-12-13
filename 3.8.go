package main

import (
	"fmt"
	"reflect"
)

func main() {
	A := []int{5, 6, 7, 8}
	B := []int{5, 6, 7, 8}
	fmt.Println(reflect.DeepEqual(A, B))
	C := []string{"Thank YOU", "SUNDAY"}
	D := []string{"SUNDAY", "THANK YOU"}
	names := []string{}
	names = append(names, "SHINOBU")
	fmt.Println(names)
	fmt.Println(reflect.DeepEqual(C, D))
}