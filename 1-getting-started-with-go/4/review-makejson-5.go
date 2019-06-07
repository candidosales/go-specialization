package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {

	consoleReader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter name: ")
	name, _ := consoleReader.ReadString('\n')

	fmt.Print("Enter address: ")
	addr, _ := consoleReader.ReadString('\n')

	db := map[string]string {
		"name": strings.TrimSpace(name),
		"address": strings.TrimSpace(addr) }

	barr, err := json.Marshal(db)

	if err != nil {
		fmt.Println("Unable to create json objects from input arguments")
	} else {
		fmt.Printf("%s", barr)
	}
}
