package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(numbers *[]int, i int) {
	(*numbers)[i+1], (*numbers)[i] = (*numbers)[i], (*numbers)[i+1]
}

func BubbleSort(numbers *[]int) {
	var (
		n      = len(*numbers)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if (*numbers)[i] > (*numbers)[i+1] {
				Swap(numbers, i)
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	numbers := make([]int, 0, 10)

	for scanner.Scan() {
		str := scanner.Text()

		if len(str) == 0 {
			break
		}

		strs := strings.Split(str, " ")
		if len(strs) < 10 {
			for i := range strs {
				intNum, _ := strconv.Atoi(strs[i])
				if len(strs[i]) != 0 {
					numbers = append(numbers, intNum)
				}
			}

			BubbleSort(&numbers)

			fmt.Println(numbers)
			break
		} else {
			fmt.Print("> 10 numbers")
			break
		}
	}

	if err := scanner.Err(); err != nil {
		os.Exit(1)
	}
}
