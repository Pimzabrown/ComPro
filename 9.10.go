package main

import "fmt"

func main() {
	n, e := fmt.Print("Hello", "PIM", 2211, 1999, "\n")
	fmt.Print("number of bytes written :", n, "\n")
	fmt.Print("write error encountered :", e, "\n")

	
}