package main

import (
	"fmt"
	"math"
)

func main() {
	var number float64
	i := 0
	for i < 2 {
		fmt.Printf("What's the number? ")
		fmt.Scan(&number)
		fmt.Println(math.Trunc(number))

		i++
	}
}
