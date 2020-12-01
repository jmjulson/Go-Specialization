package main

//Students will receive five points if the program
//sorts the integers and five
//additional points if there are four goroutines that each
//printsout a set of array elements that it is storing.
//prompt user to enter a series of ints
//array partioned into 4 go routines

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	c := make(chan []int)
	fab := make([]int,0)
	
	for i := 0; i < 4; i++ {
		fmt.Println(" Main: Starting worker", i)
		wg.Add(1)
		go runner(generateRandNums(), c, &wg, i)
		
		wg.Wait()
		w := <-c
		
		sort.Ints(w)
		for _, ww := range w{
		 fab = append(fab, ww)
		}
	}
	wg.Wait()

	sort.Ints(fab)
	fmt.Println(fab)
}

func generateRandNums() []int{
	var s []int
	for i := 0; i <=3; i++ {
		fab := rand.Intn(100)
		s = append(s, fab)
	}

	fmt.Println(s)
	return s
}

func runner(bar []int, c chan []int, wg *sync.WaitGroup, id int){
	fmt.Printf("Worker %v: Started\n", id)
	sort.Ints(bar)
	wg.Done()
	c <- bar	
}
