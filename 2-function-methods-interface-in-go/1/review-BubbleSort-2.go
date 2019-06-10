package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(data []int, index int) {
	tmp := data[index]
	data[index] = data[index+1]
	data[index+1] = tmp
}

func BubbleSort(data []int) {
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				Swap(data, j)
			}
		}
	}
}

func main() {
	data := make([]int, 0, 10)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Please, enter up to 10 integers (space separated):\n")
	scanner.Scan()
	line := scanner.Text()
	tokens := strings.Fields(line)

	for i, token := range tokens {
		if i == 10 {
			break
		}
		num, err := strconv.Atoi(token)
		if err == nil {
			data = append(data, num)
		}
	}

	BubbleSort(data)
	fmt.Printf("\nData after sorting:\n")
	fmt.Println(data)

}
