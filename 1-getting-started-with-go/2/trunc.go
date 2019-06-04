package main

import "fmt"

func main() {
	var number float32
	fmt.Printf("What's the number? ")
	fmt.Scan(&number)
	fmt.Printf("%.0f \n", number)
}
