package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	eat   string
	move  string
	speak string
}

func (a *Animal) InitMe(eat, move, speak string) {
	a.eat = eat
	a.move = move
	a.speak = speak
}

func (a *Animal) Eat() {
	fmt.Print(a.eat)
}

func (a *Animal) Move() {
	fmt.Print(a.move)
}

func (a *Animal) Speak() {
	fmt.Print(a.speak)
}

func (a *Animal) PrintInfo(info string) {
	switch info {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	}
}

func main() {
	console := bufio.NewScanner(os.Stdin)
	fmt.Print(">")
	for console.Scan() {
		str := console.Text()

		if strings.ToLower(str) == "x" {
			Exit()
		}

		inputs := strings.Fields(str)

		var animal Animal

		switch inputs[0] {
		case "cow":
			animal.InitMe("grass", "walk", "moo")
		case "bird":
			animal.InitMe("worms", "fly", "peep")
		case "snake":
			animal.InitMe("mice", "slither", "hsss")
		}
		animal.PrintInfo(inputs[1])

		fmt.Println("")
		fmt.Print(">")
	}

	if err := console.Err(); err != nil {
		Exit()
	}
}

func Exit() {
	os.Exit(1)
}
