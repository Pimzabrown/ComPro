package main

import "fmt"

func main(){
	elements := make(map[string]string)
	elements["Na"] = "Sodium"
	elements["Ca"] = "Calcium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"

	b := elements["B"]
	fmt.Println(b)

	b, found := elements["B"]
	fmt.Println(b)
	fmt.Println(found)
}