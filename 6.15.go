package main

import "fmt"

func main() {
	defer fmt.Println("HI BAM")
	var a map[int] int
	a[0] = 1
	fmt.Println(a)
}