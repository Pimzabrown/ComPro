package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println("REPLAY")
		continue
		fmt.Println("RING DING DONG")
	}
	fmt.Println("ODE EYE")
}