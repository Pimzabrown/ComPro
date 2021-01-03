package main

import "fmt"

type student struct {
	name string
	age  int
}

func (std student) introduce() {
	fmt.Println("Hello my name is", std.name)
}

type pupil struct {
	address string
	std     student
}

func main() {
	Bam := student{name: "Bam"}
	pup := pupil{std: bam}
	pup.std.introduce() 
}