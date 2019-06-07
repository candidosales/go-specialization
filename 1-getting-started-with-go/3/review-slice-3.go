package main

import (
	"fmt"
	"strconv"
    "sort"
)

func main() {
    
	var in string
	sli := make([]int, 3)
    j := 0
	for {
        fmt.Print("Please input an interger for slice(enter X to exit): ")
		fmt.Scanln(&in)
		if in == "X" {
			break
		}
		in,_ := strconv.Atoi(in)
        if j < len(sli) {   // cap() here changes from 3 to 6 when 4th element added
             sli[j] = in
        } else {
            sli = append(sli, in)   
        }
        sortedSli := make([]int, len(sli))
        copy(sortedSli,sli)
        sort.Ints(sortedSli)
        fmt.Println("sorted slice: ", sortedSli)
        j++
	}
    sortedSli := make([]int, len(sli))
    copy(sortedSli,sli)
    sort.Ints(sortedSli)
	fmt.Println(sortedSli)
}    
/*
 go run slice.go
Please input an interger for slice(enter X to exit): -10
sorted slice:  [-10 0 0]
Please input an interger for slice(enter X to exit): 20
sorted slice:  [-10 0 20]
Please input an interger for slice(enter X to exit): 3
sorted slice:  [-10 3 20]
Please input an interger for slice(enter X to exit): 4
sorted slice:  [-10 3 4 20]
Please input an interger for slice(enter X to exit): -2
sorted slice:  [-10 -2 3 4 20]
Please input an interger for slice(enter X to exit): 2
sorted slice:  [-10 -2 2 3 4 20]
Please input an interger for slice(enter X to exit): 5
sorted slice:  [-10 -2 2 3 4 5 20]
Please input an interger for slice(enter X to exit): X
[-10 -2 2 3 4 5 20]
*/