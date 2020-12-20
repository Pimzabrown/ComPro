package main

import "fmt"

func main() {
	for i := 0; i < 150; i++ {
		fmt.Println(i)
	}
	for i := 0; i > 50; i++ {
		fmt.Print(i)
	}
}