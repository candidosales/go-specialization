package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {

	fmt.Print("Enter file name: ")
	consoleReader := bufio.NewReader(os.Stdin)
	filename, _ := consoleReader.ReadString('\n')

	fd, err := os.Open(strings.TrimSpace(filename))
	if err != nil {
		fmt.Printf("Cannot locate file %s\n", filename)
	} else {

		slice := make([]*name, 0, 0)
		eof := false
		var firstName, lastName string

		for {
			firstName, lastName, eof = readLine(fd)
			newName := new(name)
			newName.fname = firstName
			newName.lname = lastName
			slice = append(slice, newName)
			if eof {	//in case of end of file break
				break
			}
		}

		fmt.Println()
		for _, v := range slice {
			fmt.Printf("%s %s\n", v.fname, v.lname)
		}
	}
}

/**
Reads a line from file and returns two fields of string (max 20 chars.) and eof bool - true in case of end of file
 */
func readLine(fd *os.File) (string, string, bool) {

	b := make([]byte, 1, 1)
	firstName:= make([]byte, 20, 20)
	lastName := make([]byte, 20, 20)
	changeField := false
	eof := false

	i := 0
	for ; ; i++ {

		_, err := fd.Read(b)

		if err != nil {	// in case of error exit
			if err.Error() == "EOF" { // end of file
				eof = true
				break
			} else {
				fmt.Println("Unable to read from file")
				os.Exit(1)
			}

		} else if b[0] == ' ' {	// in case of end of field
			changeField = true
			i = 0
		} else if b[0] == '\n' {	//in case of end of line
			break
		}

		if i < 20 {	// for fields greater than 20 characters truncate
			if changeField {
				lastName[i] = b[0]
			} else {
				firstName[i] = b[0]
			}
		}
	}

	return string(firstName), string(lastName), eof
}
