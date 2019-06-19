package main

import (
   "fmt"
   "sync"
   "bufio"
   "strconv"
   "os"
   "strings"
   _"io"
   "sort"
   _"math"
)

func main(){
	fmt.Println("Please enter integer numbers to sort.")
	fmt.Println("Kindly separate them using a space.")
	fmt.Printf("> ")
   
   scanner := bufio.NewScanner(os.Stdin)
   scanner.Scan()
   requestArray := strings.Fields(scanner.Text())

   inputInts := make([]int,0,1)
   for _,s := range requestArray{
   	i, err := strconv.Atoi(s)
   	if err == nil{
   		inputInts = append(inputInts, i)
   	}
   	//else{
   	//	io.Writer(os.Stderr, err)
   	// }
   }
   fmt.Printf("Input Array: %v\n", inputInts)
   
   arrayLength := len(inputInts)
   sliceLength := arrayLength / 4
   // make 4 non-overlapping slices to sort concurrently
   slice1 := inputInts[0:sliceLength]
   slice2 := inputInts[sliceLength : 2*sliceLength]
   slice3 := inputInts[2*sliceLength : 3*sliceLength]
   slice4 := inputInts[3*sliceLength : arrayLength]

   var wg sync.WaitGroup
   wg.Add(4)
   go sortPart(slice1, &wg)
   go sortPart(slice2, &wg)
   go sortPart(slice3, &wg)
   go sortPart(slice4, &wg)
   wg.Wait()

   // merge the 4 sorted slices in place
   fmt.Printf("Before merging: %v\n", inputInts)

   var pointerToMinIndex *int
   var currentMin int
   finalSlice := make([]int,0,arrayLength)

   // indices to first elements of each of the slices to merge
   i:=0
   j:=sliceLength
   k:=2*sliceLength
   l:=3*sliceLength
   maxInt := 1 << 32

   for i<sliceLength || j < 2*sliceLength || k < 3*sliceLength || l < arrayLength{
   	if i < sliceLength {
   		currentMin = inputInts[i]
   		pointerToMinIndex = &i
   	}

   	if j < 2 * sliceLength {
   		if inputInts[j] < currentMin{
   			currentMin = inputInts[j]
   			pointerToMinIndex = &j
   		}
   	}

   	if k < 3 * sliceLength {
   		if inputInts[k] < currentMin{
   			currentMin = inputInts[k]
   			pointerToMinIndex = &k
   		}
   	}

   	if l < arrayLength {
   		if inputInts[l] < currentMin{
   			currentMin = inputInts[l]
   			pointerToMinIndex = &l
   		}
   	}

   	// copy min to final slice and increment index of slice
   	finalSlice = append(finalSlice, currentMin)
   	*pointerToMinIndex = *pointerToMinIndex + 1
   	currentMin = maxInt
   }
   fmt.Printf("After Sorting: %v\n", finalSlice)
}

func sortPart(sli []int, wg *sync.WaitGroup){
	fmt.Printf("Will sort %v\n", sli)
   sort.Ints(sli)
   wg.Done()
}