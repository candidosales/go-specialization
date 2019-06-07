package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	var nameInput string
	var addressInput string
	var err error
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter name")
	nameInput, _ = reader.ReadString('\n')
	nameInput = strings.Replace(nameInput, "\n", "", -1)

	fmt.Println("Please enter address")
	addressInput, _ = reader.ReadString('\n')
	addressInput = strings.Replace(addressInput, "\n", "", -1)

	details := make(map[string]interface{})
	details["name"] = nameInput
	details["address"] = addressInput
	marshalledByte, err := json.Marshal(details)

	if err == nil {
		fmt.Printf("The jsonObject is %v\n", string(marshalledByte))
	} else {
		fmt.Println("Error Encountered")
	}
}
