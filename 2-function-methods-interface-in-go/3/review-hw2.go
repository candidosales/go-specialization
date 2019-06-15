package main

import "fmt"

type Animal struct {
	food string
	locomotion string
	noise string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {

	cow := Animal {"grass", "walk", "moo"}
	bird := Animal {"worms", "fly", "peep"}
	snake := Animal {"mice", "slither", "hsss"}



	for ;; {
		fmt.Print("Enter > ")
		var name, info string
		fmt.Scan(&name, &info)

		var a Animal
		switch name {
		case "snake":
			a = snake
		case "cow":
			a = cow
		case "bird":
			a = bird
		}

		switch info {
		case "eat":
			a.Eat()
		case "move":
			a.Move()
		case "speak":
			a.Speak()
		}
	}


}
