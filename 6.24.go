package main

import (
	"fmt"
)

type myError struct {
	error string
}

func (e myError) Error() string {
	if word == "Hi" {
		return myError{"Can not say Hi"}
	}
	return nil
}
func main() {
	e1 := say("Hello")
	fmt.Println(e1)
	e2 := say("HI :::::::::")
	fmt.Println(e2)
}