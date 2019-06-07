package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// function that can read a line from stdin
func readline() (l string, e error) {
	reader := bufio.NewReader(os.Stdin)
	return reader.ReadString('\n')
}

func main() {

	pii := make(map[string]string)

	fmt.Print("Enter your name: ")
	n, e := readline()
	if e != nil {
		fmt.Printf("Error %v\n", &e)
		os.Exit(1)
	}
	pii["name"] = n[:len(n)-1]
	fmt.Print("Enter your address: ")
	a, e := readline()
	if e != nil {
		fmt.Printf("Error %v\n", &e)
		os.Exit(1)
	}
	pii["address"] = a[:len(a)-1]

	fmt.Printf("%s\n%s\n", pii["name"], pii["address"])

	b, e := json.Marshal(&pii)
	if e != nil {
		fmt.Printf("Error %v\n", &e)
		os.Exit(1)
	}
	j := string(b)
	fmt.Printf("%s\n", j)
}
