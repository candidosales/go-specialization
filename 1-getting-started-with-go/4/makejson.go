package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func main() {
	var nm string
	fmt.Printf("What's the name? ")
	fmt.Scan(&nm)

	var addr string
	fmt.Printf("What's the address? ")
	fmt.Scan(&addr)

	person := Person{Name: nm, Address: addr}

	personBlob, err := json.Marshal(person)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(personBlob))
}
