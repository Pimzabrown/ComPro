package main

import "fmt"

func main() {
	name := make(map[string]string)
	name["G"] = "GIYU"
	name["S"] = "SHINOBU"
	name["T"] = "TANJIRO"
	name["N"] = "NEZUKO"
	name["I"] = "INOSUKE"
	name["Z"] = "ZENITSU"
	name["K"] = "KANAO"
	name["K"] = "KYOJURO"
	fmt.Println(name)

	delete(name, "K")
	fmt.Println(name)
}