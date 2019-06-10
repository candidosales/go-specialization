package main

import "fmt"

func main() {

	fmt.Println("enter values for acceleration, initial velocity, and initial displacement:")
	var acceleration, velocity, displacement float64
	fmt.Scan(&acceleration, &velocity, &displacement)

	fn := GenDisplaceFn(acceleration, velocity, displacement)

	fmt.Println("enter a value for time:")
	var time float64
	fmt.Scan(&time)
	fmt.Println(fn(time))


}

func GenDisplaceFn(acceleration float64, velocity float64, displacement float64) func(t float64) float64 {
	fmt.Println("acceleration: ", acceleration, " velocity: ", velocity, " displacement: ", displacement)
	return func(t float64) float64 {
		return 0.5 * acceleration * t * t + velocity * t + displacement
	}
}
