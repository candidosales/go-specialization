/*
Write a Bubble Sort program in Go. The program should prompt the user to type in
a sequence of up to 10 integers. The program should print the integers out on one
line, in sorted order, from least to greatest.

As part of this program, you should write a function called BubbleSort() which
takes a slice of integers as an argument and returns nothing. The BubbleSort()
function should modify the slice so that the elements are in sorted order.


*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Prompt the user for a string containing 10 integers, separated by space. We
// will handle parsing later.
func GetElementsFromPrompt() (input string) {
	fmt.Print("Enter up to 10 integers, separated by space: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()

	return
}

// Given a string as input, convert it to integer.
//
// This is a quick abstraction around strings.Atoi to clean up the parsing
// process.
func StringToInt(input string) (output int) {
	output, err := strconv.Atoi(input)

	if err != nil {
		log.Fatal(err)
	}

	return
}

// Given an input string, convert it to a slice of integers. Only the first 10
// elements are considered.
func ConvertInputToSlice(input string) []int {
	inputSlice := make([]int, 0, 10)

	splitInput := strings.Split(input, " ")

	for i, v := range splitInput {
		if i > 10 {
			fmt.Println("More than 10 elements entered. Truncating...")
			break
		}

		inputSlice = append(inputSlice, StringToInt(v))
	}

	return inputSlice
}

// Given a slice and a low index, perform an inplace swap for bubblesort.
func Swap(slice []int, idx int) {
	slice[idx+1], slice[idx] = slice[idx], slice[idx+1]
}

// Sort an array using the bubble sort algorithm. Bubble sort has a O(n^2) time
// complexity, but performs operations in place (O(1) space complexity).
func BubbleSort(slice []int) {
	isSorted := false

	for {
		if isSorted {
			break
		}
		isSorted = true
		for i := 0; i < len(slice)-1; i++ {
			if slice[i] > slice[i+1] {
				Swap(slice, i)
				isSorted = false
			}
		}
	}
}

// Using string builder, write a slice of integers to stdout
func WriteOutput(slice []int) {
	var b strings.Builder

	for _, v := range slice {
		fmt.Fprintf(&b, "%d ", v)
	}

	fmt.Println(b.String())
}

func main() {
	input := GetElementsFromPrompt()
	slice := ConvertInputToSlice(input)
	BubbleSort(slice)
	WriteOutput(slice)
}
