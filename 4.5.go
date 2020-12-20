package main

import "fmt"

func main() {
	r := [] bool {true, true, false, true, false, false, true, false, false}
	for _, v := range r {
		fmt.Printf("%t\n", v)
	}
}