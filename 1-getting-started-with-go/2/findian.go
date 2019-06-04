package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Printf("What's the sentence? ")
	fmt.Scan(&input)

	inputArray := prepareString(input)

	switch {
	case inputArray[0] == "i" && strings.Contains(input, "a") && inputArray[len(input)-1] == "n":
		fmt.Printf("Found! \n")
	default:
		fmt.Print("Not Found! \n")
	}
}

func prepareString(input string) []string {
	return strings.Split(strings.ToLower(strings.TrimRight(input, "\n")), "")
}
