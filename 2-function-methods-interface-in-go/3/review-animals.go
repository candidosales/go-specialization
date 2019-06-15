package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Struct to model the animals
type Animal struct {
	food       string
	locomotion string
	noise      string
}

// Function to initialize the struct with values
func (a *Animal) InitMe(f, l, n string) {
	a.food = f
	a.locomotion = l
	a.noise = n
}

// Function to print the action "eat" for a given animal
func (a *Animal) Eat() {
	fmt.Println(a.food)
}

// Function to print the action "move" for a given animal
func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

// Function to print the action "speak" for a given animal
func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

// Function to process the specific action for a given animal
func processAction(a Animal, action string) {
	switch strings.ToLower(action) {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	default:
		fmt.Println("Unknown action:", action)
	}
}

// Main function
func main() {

	// String to store the user's input
	var userInput string

	// Friendly usage message
	fmt.Println("\nPlease, insert two words separated by space in the format:\n\tanimal action\nAnimals: cow bird snake\nActions: eat move speak\n")

	// Infinite loop
	for {
		// Prompt
		fmt.Print("> ")
		// Read input from user, both words (animal action) in a single line
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			userInput = scanner.Text()
		}
		words := strings.Fields(userInput)

		// Check if there are just two words (animal and action)
		if len(words) != 2 {
			fmt.Println("Error. You must input exactly two words, animal and action.")
		} else {
			// Define the animal
			var animal Animal

			// Initialize the animal with the corresponding values and process its action
			switch strings.ToLower(words[0]) {
			case "cow":
				animal.InitMe("grass", "walk", "moo")
				processAction(animal, words[1])
			case "bird":
				animal.InitMe("worms", "fly", "peep")
				processAction(animal, words[1])
			case "snake":
				animal.InitMe("mice", "slither", "hsss")
				processAction(animal, words[1])
			default:
				fmt.Println("Unknown animal:", words[0])
			}
		}
	}
}
