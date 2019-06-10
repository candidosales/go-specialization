/*
This program prompts a user for acceleration, initial velocity, initial
displacement and time. It then computes the displacement after the entered time.
*/

package main

import (
	"fmt"
	"log"
	"math"
)

// Calculate displacement as a function of acceleration, initial velocity, initial
// displacement, and time.

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return (.5 * a * math.Pow(t, 2)) + (v0 * t) + s0
	}
}

func main() {
	var acceleration float64
	var initialVelocity float64
	var initialDisplacement float64
	var time float64
	var err error

	fmt.Print("Enter an acceleration: ")
	_, err = fmt.Scan(&acceleration)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter an initial velocity: ")
	_, err = fmt.Scan(&initialVelocity)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter an initial displacement: ")
	_, err = fmt.Scan(&initialDisplacement)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter an time: ")
	_, err = fmt.Scan(&time)

	if err != nil {
		log.Fatal(err)
	}

	// create a first class function with known params
	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	// print displacement based on fixed params and time
	fmt.Println(fn(time))
}
