package main

import "fmt"

func say() {
	fmt.Println("NOVEMBER")
}

func main() {
	defer say()
	fmt.Println("HI BAM")
}