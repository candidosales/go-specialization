package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		if strings.ToLower(string(str[0])) == "i" && strings.ToLower(string(str[len(str)-1:])) == "n" && strings.Contains(string(str), "a") {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not found!")
		}
	}
	if err := scanner.Err(); err != nil {
		os.Exit(1)
	}
}
