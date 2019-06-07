package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var name string
	var address string
	fmt.Println("Enter name: ")
	fmt.Scan(&name)
	fmt.Println("Enter address: ")
	fmt.Scan(&address)

	m := make(map[string]string)
	m["name"] = name
	m["address"] = address

	fmt.Println("map:", m)

	mapData, e := json.Marshal(m)
	if e != nil {
		fmt.Println("error : ", e)
	}
	fmt.Println(string(mapData))


}
