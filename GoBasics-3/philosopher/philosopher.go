package main

// each stick is a mutex
// each philospher gets a goroutine
import (
	"fmt"
	"sync"
	//"math/rand"
)

var wg sync.WaitGroup
type ChopS struct{ sync.Mutex }

type Philo struct {
	theDude int
	leftCS, rightCS *ChopS
}

func (p Philo) eat() {
	defer wg.Done()
	for ii := 0; ii < 3; ii++{
		p.leftCS.Lock()
		p.rightCS.Lock()
		output("starting to eat: ", p.theDude)

		p.leftCS.Unlock()
		p.rightCS.Unlock()
		output("stuffed", p.theDude)
	}
}

func output(action string, id int) {
	fmt.Printf("Philosopher #%d is %s\n", id+1, action)
}

func main() {
	count := 5
	CSticks := make([]*ChopS, 5)

	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, count)

	for i := 0; i < count; i++ {
		philos[i] = &Philo{theDude: i, leftCS: CSticks[i], rightCS: CSticks[(i+1)%count]}
		wg.Add(1)
		go philos[i].eat()
	}
	wg.Wait()
}
