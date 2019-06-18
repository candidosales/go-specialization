package main

import "fmt"

var value int

func section1(id string) {
	var s1 int
	for i := 0; i < 2; i++ {
		s1 = value
		s1++
		value = s1
		fmt.Println(id, "value:", s1)
	}
}

func section2(id string) {
	var s2 int
	for i := 0; i < 2; i++ {
		s2 = value
		s2 += 100
		value = s2
		fmt.Println(id, "value:", s2)
	}
}

func introMsg() {
	fmt.Println("\nThis program performs a race condition in Golang")
	fmt.Println("It launches two goroutines:\n\t#1 increasing a variable by units\n\t#2 increasing a variable by hundreds")
	fmt.Println("Eventually you can see that #1 contains increment of hundreds due to lack of synchronization")
	fmt.Println("in access to shared variable and non-deterministic goroutines execution ordering.")
	fmt.Println("\nLaunching goroutines. Once finished, press ENTER to exit")
}

func main() {
	value = 0

	introMsg()

	go section1("Goroutine #1 (units)")
	go section2("Goroutine #2 (hundreds)")

	fmt.Scanln()
}
