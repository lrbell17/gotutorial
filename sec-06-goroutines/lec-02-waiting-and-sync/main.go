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

	for i := range numGoRoutines {
		go producer(i + 1)
	}

	wg.Wait() // wait until the WaitGroup counter is 0

	elapsed := time.Since(start)
	fmt.Printf("main() took %v\n", elapsed)
}

func producer(id int) {

	n := (rand.Int() % 1000) + 1
	d := time.Duration(n) * time.Millisecond
	time.Sleep(d)
	fmt.Printf("Producer #%v ran for %v\n", id, d)

	wg.Done() // decrement waitgroup counter by 1
}
