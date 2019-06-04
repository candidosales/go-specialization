package main

import (
	"fmt"
)

func main() {
	var f float64
	var i int
	fmt.Printf("Enter a float number : ")
	fmt.Scanf("%f", &f)
	i = int(f)
	fmt.Printf("This your Integer!! : %d \n", i)
}
