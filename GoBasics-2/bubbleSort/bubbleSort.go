package main

import (
	"fmt"
)

func main() {
	var getValues []int = ColletInts()	
	BubbleSort(getValues)
}

func ColletInts() []int{
	sli := make([]int, 10)

	for i := 0; i < 10; i++ {
		var num = GetInts(i)
		sli[i] = num
	}

	return sli
}

func GetInts(feedback int) int {
	var foo = feedback + 1
	var entInt int

	fmt.Println("Enter a number: run",  foo, "of 10")
	
	fmt.Scanln(&entInt)
	
	return entInt
}

func BubbleSort(in []int) {
	n:= 10
	Swap(n, in)

	fmt.Println("sorted slice:", in)
}

func Swap(nn int, input []int) {
	swapped := true

	for swapped {
		swapped = false
		for i := 1; i < nn; i++ {
			if input[i-1] > input[i] {
				input[i], input[i-1] = input[i-1], input[i]
				swapped = true
			}
		}
	}
}