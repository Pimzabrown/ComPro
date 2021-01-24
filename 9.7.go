package main

import "fmt"

func main(){
	elements := make(map[string]string)
	elements["Ca"] = "Calcium"
	elements["V"] = "Vanadium"
	elements["Fe"] = "Iron"
	elements["Sn"] = "Tin"
	fmt.Println(elements)

	delete(elements, "Sn")
	fmt.Println(elements)
}