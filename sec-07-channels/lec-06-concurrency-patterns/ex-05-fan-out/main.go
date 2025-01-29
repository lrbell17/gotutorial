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
Fan-Out:  A function that launches a goroutine to consume from 1 channel and produce data on multiple channels
*/
func main() {

	done := make(chan bool)

	numbers := intGen(done) // start adding ints to channel

	s0, s1, s2 := fanOut3(numbers)
	counter(s0)
	counter(s1)
	counter(s2)

	// signal generator to exit after 1s
	time.Sleep(1 * time.Second)
	done <- true

	wg.Wait() // wait goroutines to finish
}

// splits a stream of ints into 3 separate streams
func fanOut3(in chan int) (out0, out1, out2 chan int) {

	wg.Add(1)

	out0 = make(chan int)
	out1 = make(chan int)
	out2 = make(chan int)

	go func() {
		defer wg.Done()

		for v := range in {
			select {
			case out0 <- v:
			case out1 <- v:
			case out2 <- v:
			}
		}

		close(out0)
		close(out1)
		close(out2)
	}()

	return
}

func passThroughCounter(in chan int) (out chan int) {
	wg.Add(1)
	out = make(chan int)

	go func() {
		defer wg.Done()
		log.Info("Counter (PT) - starting work")
		start := time.Now()
		var count int
		for v := range in {
			count++
			out <- v
		}
		fmt.Printf("Counter (PT) processed %v  items in %v\n", count, time.Since(start))
		close(out)
	}()
	return
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
