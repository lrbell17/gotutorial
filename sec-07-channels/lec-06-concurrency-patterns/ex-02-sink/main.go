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
Sink/Terminator/Consumer: A function that launches a goroutine to consume data on a channel
*/
func main() {

	done := make(chan bool)

	d := intGen(done) // start adding ints to channel
	counter(d)        // start sink to count ints received on the channel

	// signal generator to exit after 1s
	time.Sleep(1 * time.Second)
	done <- true

	wg.Wait() // wait goroutines to finish
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
