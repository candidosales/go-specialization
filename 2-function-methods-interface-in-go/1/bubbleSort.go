package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	console := bufio.NewScanner(os.Stdin)
	fmt.Println("Please insert the numbers separated by space:  ")

	for console.Scan() {
		str := console.Text()

		if strings.ToLower(str) == "x" {
			Exit()
		}

		slice := strings.Split(str, " ")
		sliceInt := ConvertArrayStringToInt(slice)

		BubbleSort(sliceInt)
		fmt.Println(sliceInt)
	}

	if err := console.Err(); err != nil {
		Exit()
	}
}

func Exit() {
	os.Exit(1)
}

func ConvertArrayStringToInt(slice []string) []int {
	sliceInt := make([]int, len(slice))

	for i, v := range slice {
		sliceInt[i], _ = strconv.Atoi(v)
	}

	return sliceInt
}

func BubbleSort(slice []int) {
	swapped := true
	for swapped {
		swapped = false
		n := len(slice)
		for i := 1; i < n; i++ {
			if slice[i-1] > slice[i] {
				Swap(slice, i)
				swapped = true
			}
		}
	}
}

func Swap(slice []int, index int) {
	slice[index], slice[index-1] = slice[index-1], slice[index]
}
