package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (c Cow) Eat() {
	println("grass")
}
func (c Cow) Move() {
	println("walk")
}
func (c Cow) Speak() {
	println("moo")
}

type Bird struct{}

func (b Bird) Eat() {
	println("worms")
}
func (b Bird) Move() {
	println("fly")
}
func (b Bird) Speak() {
	println("peep")
}

type Snake struct{}

func (s Snake) Eat() {
	println("mice")
}

func (s Snake) Move() {
	println("slither")
}

func (s Snake) Speak() {
	println("hsss")
}

func main() {
	animals := make(map[string]Animal)

	stdin := bufio.NewReader(os.Stdin)

	for {
		println("Print command  or query>")
		r, err := stdin.ReadString('\n')

		if err != nil {
			fmt.Printf("Error happened: %s, repeating\n", err.Error())
			continue
		}

		req := strings.Split(strings.TrimSuffix(r, "\n"), " ")
		if len(req) < 3 {
			fmt.Printf("Command or query must have 3 parameters, buw was: %d\n", len(req))
			continue
		}

		command := req[0]

		switch command {
		case "newanimal":
			name, kind := req[1], req[2]
			a, err := NewAnimal(kind)
			if err != nil {
				fmt.Println(err)
				continue
			} else {
				animals[name] = a
			}
			println("New animal added")
		case "query":
			name, prop := req[1], req[2]
			animal := animals[name]
			PrintProp(animal, prop)
		}
	}
}

func NewAnimal(kind string) (Animal, error) {
	switch kind {
	case "cow":
		return new(Cow), nil
	case "snake":
		return new(Snake), nil
	case "bird":
		return new(Bird), nil
	default:
		return nil, errors.New("unknown animal type")
	}
}

func PrintProp(animal Animal, prop string) {
	switch prop {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Printf("Unknown animal property '%s'\n", prop)
	}
}
