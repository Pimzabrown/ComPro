package main

import "fmt"

func main() {
	elements := make(map[string]map[string]string)
	elements["C"] = map[string]string{"name": "Carbon", "state": "gas"}
	elements["O"] = map[string]string{"name": "Oxygen", "state": "gas"}
	elements["He"] = map[string]string{"name": "Helium", "state": "gas"}
	elements["Ne"] = map[string]string{"name": "Neon", "state": "gas"}
	fmt.Println(elements)
	fmt.Println(elements["O"]["state"])
	fmt.Println(elements["He"]["state"])
	fmt.Println(elements["Ne"]["state"])
}