package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	// s  = rand.NewSource(time.Now().Unix())
	// r  = rand.New(s)
	wg sync.WaitGroup
)

/*
Fan-In:  A function that launches a goroutine to consume from multiple channels and produce data on one channel

Note:
- the multiple --> 1 mapping is for simplicity but is not necessarily true
- e.g we could have 10:2 mapping or any other combination
*/
func main() {

	done := make(chan bool)

	numbers := intGen(done)
	letters := alphaGen(done)

	alphNum := fanIn2(done, numbers, letters)
	printer(alphNum)

	// signal generator to exit after 1s
	time.Sleep(1 * time.Millisecond)
	done <- true
	done <- true
	done <- true
	wg.Wait() // wait goroutines to finish

}

// // Fan in from 2 channels to 1
func fanIn2(done chan bool, intCh chan int, alphaCh chan rune) (out chan string) {

	wg.Add(1)
	out = make(chan string)

	go func() {
		defer wg.Done()
		for {
			select {
			case n, ok := <-intCh:
				if ok {
					out <- fmt.Sprintf("%v", n)
				}
			case l, ok := <-alphaCh:
				if ok {
					out <- string(l)
				}
			case <-done:
				close(out)
				return
			}
		}
	}()

	return
}

// sink to print results
func printer(in chan string) {
	wg.Add(1)

	go func() {

		for v := range in {
			fmt.Printf("%v,", v)
		}
		fmt.Println()
		wg.Done()
	}()
}

func alphaGen(done chan bool) (out chan rune) {
	wg.Add(1)
	out = make(chan rune)

	go func() {
		defer wg.Done()
		letters := []rune("abcdefghijklmnopqrstuvwxzy")
		l := len(letters)
		for {
			select {
			case out <- letters[rand.Int()%l]:
			case <-done:
				close(out)
				return
			}
		}
	}()
	return
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
