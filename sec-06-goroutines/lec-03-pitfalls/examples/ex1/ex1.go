package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const numGoRoutines = 4

var wg sync.WaitGroup

func main() {

	start := time.Now()

	wg.Add(numGoRoutines) // Add delta to WaitGroup
	// wg.Add(numGoRoutines - 1) --> results in only 3/4 goroutines running
	// wg.Add(numGoRoutines + 1) --> results in deadlock (all goroutines are asleep)

	for i := range numGoRoutines {
		go producer(i + 1)
	}

	wg.Wait()

	// producer(numGoRoutines + 1) --> results in error due to negative waitgroup counter

	elapsed := time.Since(start)
	fmt.Printf("main() took %v\n", elapsed)
}

func producer(id int) {

	n := (rand.Int() % 1000) + 1
	d := time.Duration(n) * time.Millisecond
	time.Sleep(d)
	fmt.Printf("Producer #%v ran for %v\n", id, d)

	wg.Done()
}
