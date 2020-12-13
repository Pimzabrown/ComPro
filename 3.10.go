package main

import "fmt"

func main() {
	fmt.Printf("Hi! MY NAME IS %s, I AM %d YEARS OLD\n", "TANJIRO", 15)

	fmt.Printf("type = %T \n", 30122556556)
	fmt.Printf("num = %f \n",61336821)
	fmt.Printf("num = %2.2f \n", 4164455)
	fmt.Printf("num = %9.f \n", 780336986256)
	fmt.Printf("num = %-9.f \n", 78995895325565)
	fmt.Printf("num = %09.f \n", 69821405)
	fmt.Printf("num = %09.2f \n", 987896)
	fmt.Printf("num = %E \n", 74569863)

	fmt.Printf("2 > 3 = %t \n", 2 > 3)
	fmt.Printf("20 (base 6) = %b \n", 20)
	fmt.Printf("20 (base 9) = %o \n", 20)
	fmt.Printf("20 (base 14) = %d \n", 20)
	fmt.Printf("20 (base 18) = %x \n", 20)
}