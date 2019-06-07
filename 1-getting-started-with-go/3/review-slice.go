package main

import (
    "fmt"
    "sort"
    "strconv"
)

func printSlice(x []int) {
    fmt.Printf("%v\n", x)
}

func main() {
    x := make([]int, 0, 3)
    var s string
    for {
        fmt.Scan(&s)
        if s == "X" {
            break
        }
        i, err := strconv.Atoi(s)
        if err != nil {
            fmt.Printf("Invalid input.\n")
            continue
        }
        x = append(x, i)
        sort.Slice(x, func(i, j int) bool {return x[i] < x[j]})
        printSlice(x)
    }
}
