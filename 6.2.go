package main

import "fmt"

func main() {
	var a int
ReadInput :
	fmt.Print("ENTER NUMBER ::")
	fmt.Scan(&a)
	if a >= 10 {
		goto ReadInput
	}
	fmt.Println(a)
}