package main

import "fmt"

func writeLine(a ...interface{}) {
	for _, v := range a {
		fmt.Println(v)
	}
}

func main() {
	writeLine(22, 11, 1999, 2.22, 11.1, 19.99, "Hi MY name is Phimchanok My nickname is Bam" , true)
}