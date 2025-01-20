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

	wg.Add(numGoRoutines) // Add delta to WaitGroup
	// // wg.Add(numGoRoutines - 1) --> results in only 3/4 goroutines running
	// // wg.Add(numGoRoutines + 1) --> results in deadlock (all goroutines are asleep)

	for i := range numGoRoutines {
		go producer(i + 1)
	}
	wg.Wait()

	// producer(numGoRoutines + 1) --> results in error due to negative waitgroup counter

	for i := range numGoRoutines {
		producer2(i + 1)
	}
	wg.Wait()

	// launching goroutines from anonymous functions
	launchWorkers(5)
	wg.Wait()

	// passing waitgroup to function
	for i := range numGoRoutines {
		producer3(wg, i+1)
	}
	wg.Wait() // we won't actually wait since we are just passing a copy to our function
}

func producer(id int) {

	n := (rand.Int() % 1000) + 1
	d := time.Duration(n) * time.Millisecond
	time.Sleep(d)
	fmt.Printf("Producer #%v ran for %v\n", id, d)

	wg.Done()
}

// Using waitgroups within the function
func producer2(id int) {

	wg.Add(1)
	go func() {

		n := (rand.Int() % 1000) + 1
		d := time.Duration(n) * time.Millisecond
		time.Sleep(d)
		fmt.Printf("Producer2 #%v ran for %v\n", id, d)

		wg.Done()
	}()
}

// Referencing variables
func launchWorkers(n int) {
	for i := 0; i < n; i++ { // can't capture i because it's a common varaible
		wg.Add(1)
		id := i // need to have new variable created every iteration
		go func() {
			fmt.Printf("I am worker %v\n", id)
			wg.Done()
		}()
	}
}

// Passing waitgroups to the function
func producer3(wg sync.WaitGroup, id int) {

	wg.Add(1)
	go func() {

		n := (rand.Int() % 1000) + 1
		d := time.Duration(n) * time.Millisecond
		time.Sleep(d)
		fmt.Printf("Producer3 #%v ran for %v\n", id, d)

		wg.Done()
	}()
}
