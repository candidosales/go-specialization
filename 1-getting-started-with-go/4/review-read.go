package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Name structure initialization
type Name struct {
	fname string
	lname string
}

func main() {
	NameS := make([]Name, 0, 3)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please enter the file name: ")
	scanner.Scan()
	inputName := scanner.Text()

	//file, err := ioutil.ReadFile(inputName + ".txt")
	file, err := os.Open(inputName + ".txt")
	if err != nil {
		fmt.Print(err)
	}

	A := bufio.NewScanner(file)

	for A.Scan() {
		sli := strings.Split(A.Text(), " ")
		var person Name
		person.fname, person.lname = sli[0], sli[1]
		NameS = append(NameS, person)
	}

	for _, i := range NameS {
		fmt.Println(i.fname, i.lname)
	}
}
