package main

import (
	"fmt"
	"sort"
	"sync"
)

var wg sync.WaitGroup

func mergeSort(numbers []int) {
	fmt.Printf("Sorting array: %v\n", numbers)
	sort.Ints(numbers)
	wg.Done()
}

func merge(numbers []int) {
	middle := len(numbers) / 2
	helper := make([]int, len(numbers))
	copy(helper, numbers)

	helperLeft := 0
	helperRight := middle
	current := 0
	high := len(numbers) - 1

	for helperLeft <= middle-1 && helperRight <= high {
		if helper[helperLeft] <= helper[helperRight] {
			numbers[current] = helper[helperLeft]
			helperLeft++
		} else {
			numbers[current] = helper[helperRight]
			helperRight++
		}
		current++
	}

	for helperLeft <= middle-1 {
		numbers[current] = helper[helperLeft]
		current++
		helperLeft++
	}

	for helperRight <= high {
		numbers[current] = helper[helperRight]
		current++
		helperRight++
	}
	wg.Done()
}

func parallelMergeSort(numbers []int) {
	size := len(numbers)
	wg.Add(4)
	go mergeSort(numbers[:size/4])
	go mergeSort(numbers[size/4 : size/2])
	go mergeSort(numbers[size/2 : 3*size/4])
	go mergeSort(numbers[3*size/4:])
	wg.Wait()
	wg.Add(2)
	go merge(numbers[:size/2])
	go merge(numbers[size/2:])
	wg.Wait()
	wg.Add(1)
	merge(numbers)
	wg.Wait()
}

func main() {
	var input int
	inputs := make([]int, 0)

	fmt.Println("Please enter a list of numbers, enter x to end input.")
	for {
		_, err := fmt.Scan(&input)
		if err != nil {
			break
		}
		inputs = append(inputs, input)
	}
	parallelMergeSort(inputs)
	fmt.Println(inputs)
}
