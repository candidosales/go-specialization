package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, address string
	list := map[string]string{}

	fmt.Println("Name: ")
	fmt.Scanln(&name)
	fmt.Println("Address: ")
	fmt.Scanln(&address)
	list[name] = address

	jsonData, _ := json.Marshal(list)

	fmt.Println(string(jsonData))
}
