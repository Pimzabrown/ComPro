package main

import "fmt"

func main() {
	for i := 1; i <= 100; i = i + 1 {
		if i%2 == 0 {
			fmt.Println(i, "WINDY")
		} else {
			fmt.Println(i, "RAINNY")
		}
	}
}
