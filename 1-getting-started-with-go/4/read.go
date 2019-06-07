package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	sli := make([]Name, 0)
	// filePath := "/file.txt"

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please enter the file name: ")
	scanner.Scan()
	filePath := scanner.Text()

	file, err := os.Open(filePath + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		text := fileScanner.Text()
		fmt.Printf("Read line: %s\n", text)
		sli = append(sli, readName(text))
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, n := range sli {
		fmt.Printf("First Name: %s, Last Name: %s \n", n.fname, n.lname)
	}
}

func readName(text string) Name {
	split := strings.Split(text, " ")
	return Name{fname: split[0], lname: split[1]}
}
