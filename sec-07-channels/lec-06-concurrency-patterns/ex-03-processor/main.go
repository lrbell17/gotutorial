package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

var wg sync.WaitGroup

/*
Processor: A function that launches a goroutine to consume and produce data
*/
func main() {

	done := make(chan bool)

	d := intGen(done) // start adding ints to channel

	// filter odd numbers
	d = filter(d, func(x int) bool {
		return x%2 == 1
	})

	counter(d) // start sink to count ints received on the channel

	// signal generator to exit after 1s
	time.Sleep(1 * time.Second)
	done <- true

	wg.Wait() // wait goroutines to finish
}

// filter keeps numbers matching condition based on predicate functions
func filter(in chan int, f func(x int) bool) (out chan int) {
	wg.Add(1)
	out = make(chan int)

	if f == nil {
		f = func(int) bool { return true }
	}

	go func() {
		defer wg.Done()
		for v := range in {

			// add to out channel if predicate function is satisfied
			if f(v) {
				out <- v
			}
		}
		close(out)
	}()

	return
}

// sink to count ints received on the channel
func counter(in chan int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Info("Counter - starting work")
		start := time.Now()
		var count int
		for range in {
			count++
		}
		fmt.Printf("Counter processed %v items in %v\n", count, time.Since(start))
	}()
}

// random number generator - returns channel on which random ints are produced
func intGen(done chan bool) (out chan int) {
	wg.Add(1)
	out = make(chan int)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				close(out)
				return
			case out <- rand.Int() % 200:
			}
		}
	}()
	return out
}
