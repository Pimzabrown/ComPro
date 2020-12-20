package main

import "fmt"

func say() {
	fmt.Println("You had me at hello")
}
func greet(name string) {
	fmt.Println("Hello", name)
}
func main() {
	say()
	greet("MARKLEE")
}