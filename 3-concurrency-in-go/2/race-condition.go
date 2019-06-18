package main

import (
	"fmt"
)

/*
Here, we can see that the getNumber function is setting the value of i in a separate goroutine.
We are also returning i from the function without any knowledge of whether our goroutine has completed
 or not. So now, there are two operations that are taking place:

1. The value of i is being set to 5
2. The value of i is being returned from the function

Now depending on which of these two operations completes first, our value printed will be either 0 (the default integer value) or 5.

This is why it’s called a data race : the value returned from getNumber changes depending on which of the operations 1 or 2 finish first.

*/

func main() {
	fmt.Println(getNumber())
}

func getNumber() int {
	var i int
	go func() {
		i = 5
	}()

	return i
}
