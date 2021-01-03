package main

import "fmt"

type student struct {
	name string
	age int
	email string
}

func main() {
	std := student{name: "BAM"}
	p := &std
	(*p).age = 20
	p.email = "BAMMMM@gmail.com"

	fmt.Println(std)
}