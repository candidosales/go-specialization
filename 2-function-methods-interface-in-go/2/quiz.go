// package main

// import (
// 	"fmt"
// )

// func fA() func() int {
// 	i := 0
// 	return func() int {
// 		i++
// 		return i
// 	}
// }

// func main() {

// 	fB := fA()
// 	fmt.Print(fB())
// 	fmt.Print(fB())

// }

package main

import "fmt"

func main() {

	i := 1

	fmt.Print(i)

	i++

	defer fmt.Print(i + 1)

	fmt.Print(i)

}


1. Yes
2. func fA(fB func (int) string)
3. A function with no name
4. func fA(fB (int) string) func (string) int / func fA(int) string {} / func fA(fB func (int) string) {} (falra usat) / func fA() fB func (string) int{} - last
5. 12
6. "..."
7. 123