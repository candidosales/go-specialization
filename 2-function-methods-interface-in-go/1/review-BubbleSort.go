package main

import "fmt"

func main() {


	s1  := []int{}
	len := 0
	for i := 0; i < 10; i++ {
		fmt.Println("Enter number: ")
		var num int
		fmt.Scan(&num)
		s1 = append(s1, num)
		len++
	}

	fmt.Println(s1)

	BubbleSort(s1[0: len + 1])
	fmt.Println(s1)
}

func BubbleSort(s []int) {
	for j:=0;j<len(s);j++{
		for i:=0;i<len(s)-1;i++ {
			Swap(s,i)
		}
	}


}

func Swap(s []int, i int) {
	if s[i] > s[i + 1] {
		s[i], s[i+1] = s[i+1], s[i]
	}
}
