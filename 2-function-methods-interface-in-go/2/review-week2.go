package main

import "fmt"

func GenDisplaceFn(a, v, s float64) func(t float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t+v*t+s
	}
}

func main() {
	var a, v, s, t float64

	fmt.Println("Write the acceleration: ")
	fmt.Scanln(&a)
	fmt.Println("Write the initial velocity: ")
	fmt.Scanln(&v)
	fmt.Println("Write the initial displacement: ")
	fmt.Scanln(&s)

	fmt.Println("Write the time: ")
	fmt.Scanln(&t)

	fn := GenDisplaceFn(a, v, s)
	fmt.Println(fn(t))
}
