package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("test.txt")
	fmt.Print(err)
}


1. When one thing can have multiple forms.
2. Inheritance and overriding enable polymorphism.
3. The type defines all methods specified in the interface.
4. An interface always has a dynamic type.
5. I, II, and II
6. It can be used to allow a function to accept any type as a parameter.
7. nil