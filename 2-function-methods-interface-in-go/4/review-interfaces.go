package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (cow Cow) Eat() {
	fmt.Println("grass")
}
func (cow Cow) Move() {
	fmt.Println("walk")
}
func (cow Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{}

func (bird Bird) Eat() {
	fmt.Println("worms")
}
func (bird Bird) Move() {
	fmt.Println("fly")
}
func (bird Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{}

func (snake Snake) Eat() {
	fmt.Println("mice")
}
func (snake Snake) Move() {
	fmt.Println("slither")
}
func (snake Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	var command, name, request string
	animalNameMap := map[string]Animal{}
	for {
		fmt.Print(">")
		fmt.Scan(&command)
		fmt.Scan(&name)
		fmt.Scan(&request)
		var animal Animal
		switch command {
		case "newanimal":
			switch request {
			case "cow":
				animal = Cow{}
			case "bird":
				animal = Bird{}
			case "snake":
				animal = Snake{}
			}
			animalNameMap[name] = animal
			fmt.Println("Created it!")

		case "query":
			animal = animalNameMap[name]
			switch request {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			}
		}
	}
}
