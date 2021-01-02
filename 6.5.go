package main

import "fmt"

func main() {
	var p *int
	i := 89
	fmt.Println("value", i)
	p = &i 
	*p = 5
	fmt.Println("value", i)
}