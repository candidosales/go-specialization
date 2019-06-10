package main

import "fmt"

func GenDisplaceFn(accl, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return t*t*accl*0.5 + t*v0 + s0
	}
}
func main() {
	args := make([]float64, 4)

	for i := 0; i < 4; i++ {
		var numFloat float64
		switch i {
		case 0:
			fmt.Print("Acceleration? ")
		case 1:
			fmt.Print("Initial velocity? ")
		case 2:
			fmt.Print("Input initial displacement? ")
		case 3:
			fmt.Print("Input value of time? ")
		}

		_, err := fmt.Scan(&numFloat)
		if err != nil {
			fmt.Println("please input a valid floating point number!")
			return
		}
		args[i] = numFloat
	}

	fn := GenDisplaceFn(args[0], args[1], args[2])
	fmt.Println("displacement at t:", fn(args[3]))
}
