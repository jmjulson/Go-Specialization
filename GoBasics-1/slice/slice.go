package main

import (
	"fmt"
)

func main() {
	addPromptVal()
}

func addPromptVal() []int {
	sli := make([]int, 3)
	var num int

	for i := 0; i >= 0; i++ {
		fmt.Println("Enter a number:")
		fmt.Scanln(&num)

		// update the intial 3 value - add to slice outside of :2
		if i < 3 {
			sli[i] = num
		} else {
			sli = append(sli, num)
		}

		//the array is already sorted - linear input
		fmt.Println(sli)

		//check if the user wants to exit
		check := checkStop()

		if check == "X" {
			break
		}
	}

	return sli
}

func checkStop() string {
	var foo string

	fmt.Println("Enter \"X\" to end program (note the casing), else hit Return to Continue")
	fmt.Scanln(&foo)

	return foo
}
