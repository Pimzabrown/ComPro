package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, n :=  range numbers {
		total = total + n
	}
	return total
}

func main() {
	x := sum(2, 4, 5, 8, 10, 15, 22, 2, 54, 79, 36, 85, 20, 78, 45)
	fmt.Println(x)

	y := sum()
	fmt.Println(y)

}