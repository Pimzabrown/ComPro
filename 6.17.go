package main 

import "fmt"

func main() {
	panic("MY NAME IS PHIMCHANOK")
	text := recover()
	fmt.Println(text)
}