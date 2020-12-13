package main

import "fmt"

func main() {

var p *int
i := 897
fmt.Println("value", i)
p = &i 
*p = 41
fmt.Println("value", i)
}