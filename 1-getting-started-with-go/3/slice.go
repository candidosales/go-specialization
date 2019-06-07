package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	a := make([]int, 0, 3)

	for scanner.Scan() {
		str := scanner.Text()

		if strings.ToLower(str) == "x" {
			os.Exit(1)
		}

		if value, err := strconv.Atoi(str); err == nil {
			a = append(a, value)
			sort.Ints(a)
			fmt.Println(a)
		}
	}
	if err := scanner.Err(); err != nil {
		os.Exit(1)
	}
}
