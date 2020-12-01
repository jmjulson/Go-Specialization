package main

import (
	"fmt"
)

func main() {
    wait := make(chan struct{})
    n := 0
    go func() {
		// initial routine
        n++
        close(wait)
	}()
	// the n++ introduces the race condition
    n++
	<-wait
	a := n
    fmt.Println(a)
}

// example provided by yourbasic.org