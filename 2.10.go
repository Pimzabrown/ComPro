package main

import "fmt"

var greeting = "Sunday"

func main() {
	name := "Palim"
	age := 21
	fmt.Println(greeting)
	fmt.Println("My name is", name)
	fmt.Println("I am", age, "years old")
}

func f() {
	fmt.Println(greeting)
	fmt.Println("My name is", name)
	fmt.Println("I am", age, "years old")
}