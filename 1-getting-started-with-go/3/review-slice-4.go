package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sl := make([]int, 0, 3)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		if str == "X" {
			break
		}

		num, err := strconv.Atoi(str)
		if err == nil {
			sl = append(sl, num)
			sort.Ints(sl)
		}

		fmt.Println(sl)
	}

	if err := scanner.Err(); err != nil {
		os.Exit(1)
	}
}
