package main

import (
	"fmt"
	"math/rand"
)

/*
Generator: A function that launches a goroutine to produce data on a channel it returns
*/
func main() {

	done := make(chan bool)

	d := intGen(done)
	for i := 0; i < 10; i++ {
		fmt.Println(<-d)
	}
	done <- true // signal goroutine to exit gracefully
}

// random number generator - returns channel on which random ints are produced
func intGen(done chan bool) (out chan int) {
	out = make(chan int)
	go func() {
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
