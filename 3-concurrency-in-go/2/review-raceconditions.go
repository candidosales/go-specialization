/*
Race Conditions:
A race condition is when two or more routines have access to the same resource, such as a variable or 
data structure and attempt to read and write to that resource without any regard to the other routines.
Outcome depends on non-deterministic ordering.
Races occur due to communication.

Please run the program using the flag -race so you can see the data race
go run -race raceconditions.go
*/
package main

import (
	"fmt"
	"time"
)

var i int

func main() {
	fmt.Println(getNumber())
}

func getNumber() int {
	go incrementby5()
	go incrementby10()
	time.Sleep(1000 * time.Millisecond)
	return i
}

func incrementby5() int {
	i = 5
	return i;
}

func incrementby10() int {
	i = 10
	return i;
}