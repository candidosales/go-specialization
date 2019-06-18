package main

import "fmt"

// Race condition occurs when the more than one threads are trying to access and work (make changes to) on the
// same resource (in a program, that would be a variable).

// In this program, the first goroutine adds 10 to the integer and the second one subtracts 10 from it.
// since both of these are run concurrently, the output will be uncertain

func main() {
	integer := 100

	go func(integerToBeChanged *int) {
		// add 10 to the integer
		*integerToBeChanged += 10
		fmt.Println(*integerToBeChanged)
	}(&integer)

	go func(integerToBeChanged *int) {
		// subtract 10 from integer
		*integerToBeChanged -= 10
		fmt.Println(*integerToBeChanged)
	}(&integer)

	integer += 10

	fmt.Println(integer)
}
