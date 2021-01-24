package main

import "fmt"

func main() {
	fmt.Printf("My name is %s, I am %d years old\n", "DEKU", 15 )

	fmt.Printf("type = %T \n", 3.14159265359)
	fmt.Printf("pi = %f \n", 3.1459265359)

	fmt.Printf("1 > 9 = %t \n", 1 > 9)
	fmt.Printf("10 (base 16) = %d \n", 10)
	fmt.Printf("10 (base 22) = %o \n", 10)
	fmt.Printf("10 (base 1) = %b \n", 10)
	fmt.Printf("10 (base 11) = %x \n", 10)
}