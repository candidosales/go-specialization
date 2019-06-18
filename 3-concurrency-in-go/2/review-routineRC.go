package main

import (
	"fmt"
	"time"
)

/*
An addition function. This function takes two integers and adds value of second integer to first
*/
func addCounter(x, y *int) {
	*x = *x + *y
}

/*
A multiply function. This function takes two integers and updates the first integer to product of both of the integers
*/
func multiplyCounter(x, y *int) {
	*x = *x * *y
}

/*
Topic: Write two goroutines which have a race condition when executed concurrently.
Explain what the race condition is and how it can occur.

Usage: Please run "go build routineRC.go" and run the executable 3-5 times to see the race condition in action
*/
func main() {

	// Assigning the variable x to a zero value
	x := 0

	/*
		A go routine that runs in multiplyCounter function in a for loop with enough sleep time.
		We are updating the value of x after each iteration of the loop.
		This causes the multiplyCounter for loop to execute concurrently with addCounter for loop which also updates x
	*/
	go func() {
		for j := 1; j < 9; j++ {
			multiplyCounter(&x, &j)
			time.Sleep(41 * time.Millisecond)
			fmt.Println("j =", j, "and x =", x)
		}
		// Out of 3-5 runs you would mostly find value of x at the end of multiplyCounter loop is different from the previous run(s)
		fmt.Println("The value of x after 8 iterations of multiplyCounter is", x)
	}()

	/*
		This the addCounter for loop.
		We are also updating the value of x after each iteration of this loop.
		This causes communication condition between both the loops for value of x and hence introduces a race condition.
	*/
	for i := 1; i < 9; i++ {
		addCounter(&x, &i)
		time.Sleep(42 * time.Millisecond)
		fmt.Println("i =", i, "and x =", x)
	}

	// Out of 3-5 runs you would mostly find value of x at the end of addCounter loop is different from the previous run(s)
	fmt.Println("The value of x after 8 iterations of addCounter is", x)

	// This defer function keeps the go routine alive in case addCounter finishes.
	defer func() {
		fmt.Println("Press enter/return key to exit")
		fmt.Scanln()
	}()
}
