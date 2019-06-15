package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*---------------------------------------------------
	Interface "Animal" and function "ProcessAction"
  ---------------------------------------------------*/

// Interface for animals
type Animal interface {
	Eat()
	Move()
	Speak()
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

/*-------
	Cow
  -------*/

// Struct to model the Cow
type Cow struct {
	food       string
	locomotion string
	noise      string
}

// Function to print the action "eat" for a Cow
func (c Cow) Eat() {
	fmt.Println(c.food)
}

// Function to print the action "move" for a Cow
func (c Cow) Move() {
	fmt.Println(c.locomotion)
}

// Function to print the action "speak" for a Cow
func (c Cow) Speak() {
	fmt.Println(c.noise)
}

/*--------
	Bird
  --------*/

// Struct to model the Bird
type Bird struct {
	food       string
	locomotion string
	noise      string
}

// Function to print the action "eat" for a Bird
func (b Bird) Eat() {
	fmt.Println(b.food)
}

// Function to print the action "move" for a Bird
func (b Bird) Move() {
	fmt.Println(b.locomotion)
}

// Function to print the action "speak" for a Bird
func (b Bird) Speak() {
	fmt.Println(b.noise)
}

/*---------
	Snake
  ---------*/

// Struct to model the Snake
type Snake struct {
	food       string
	locomotion string
	noise      string
}

// Function to print the action "eat" for a Snake
func (s Snake) Eat() {
	fmt.Println(s.food)
}

// Function to print the action "move" for a Snake
func (s Snake) Move() {
	fmt.Println(s.locomotion)
}

// Function to print the action "speak" for a Snake
func (s Snake) Speak() {
	fmt.Println(s.noise)
}

/*---------------
	Main method
  ---------------*/

// Main function
func main() {

	// Map to store the animals to be created
	animals := make(map[string]Animal)

	// String to store the user's input
	var userInput string

	/*
		Each “newanimal” command must be a single line containing three strings.
			The first string is “newanimal”.
			The second string is an arbitrary string which will be the name of the new animal.
			The third string is the type of the new animal, either “cow”, “bird”, or “snake”.
		Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.

		Each “query” command must be a single line containing 3 strings.
			The first string is “query”.
			The second string is the name of the animal.
			The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
		Your program should process each query command by printing out the requested data.
	*/

	// Friendly usage message
	fmt.Println("\nPlease, insert a command. It can be <newanimal> or <query>")
	fmt.Println("\tExamples:\n\t\tnewanimal <name> <cow | bird | snake>")
	fmt.Println("\t\tquery <name> <eat | move | speak>")

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

		// Check the option (newanimal or query). It must contain three strings in the same line
		if len(words) != 3 {
			fmt.Println("Error. You must input exactly three words: command, animal name and type or action.")
		} else {

			// if newanimal, let's create the new animal
			if strings.ToLower(words[0]) == "newanimal" {

				// Check for animal presence in the map (if it has been previously created)
				_, presence := animals[words[1]]
				if presence {
					fmt.Printf("Error. The animal named %s already exist.\n", words[1])
				} else {
					// Initialize the animal with the corresponding name and type
					switch strings.ToLower(words[2]) {
					case "cow":
						animals[words[1]] = Cow{"grass", "walk", "moo"}
						fmt.Println("Created it!")
					case "bird":
						animals[words[1]] = Bird{"worms", "fly", "peep"}
						fmt.Println("Created it!")
					case "snake":
						animals[words[1]] = Snake{"mice", "slither", "hsss"}
						fmt.Println("Created it!")
					default:
						fmt.Println("Unknown animal type:", words[2])
					}
				}

				// if query, let's create the new animal
			} else if strings.ToLower(words[0]) == "query" {

				// Check for animal presence in the map (if it has been previously created)
				_, presence := animals[words[1]]
				if presence {
					// Perform appropiate action for the given animal
					switch strings.ToLower(words[2]) {
					case "eat":
						animals[words[1]].Eat()
					case "move":
						animals[words[1]].Move()
					case "speak":
						animals[words[1]].Speak()
					default:
						fmt.Println("Unknown action:", words[2])
					}
				} else {
					fmt.Printf("The animal named %s does not exist.\n", words[1])
				}

				// if unknown command
			} else {
				fmt.Println("Error. Unknown command:", words[0])
			}
		}
	}
}
