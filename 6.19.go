package main

import "fmt"

type student struct {
	name string
	age int
	email string
}

func main() {
	var std [10]student
	std[0] = student{"BAM", 20, "BAMMMM@gmail.co.th"}

	fmt.Println(std[0])
	fmt.Println(std[0].name)
}