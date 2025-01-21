package main

import (
	"fmt"
	"sync"
)

const numWorkers = 40

var (
	wg           sync.WaitGroup
	counter      = 0
	counterMutex sync.Mutex

	mapCounter      = 0
	mapCounterMutex sync.Mutex
	m               map[int]int
)

func main() {

	// Using counter
	for i := 0; i < numWorkers; i++ {
		updateCounter()
	}
	wg.Wait()

	fmt.Printf("Counter: %v\n", counter)

	// Using map
	m = make(map[int]int)
	for i := 0; i < numWorkers; i++ {
		updateCounterMap(i)
	}
	wg.Wait()
	fmt.Printf("Map: %v\n, Counter: %v\n", m, mapCounter)

}

func updateCounter() {
	wg.Add(1)
	go func() {
		/*
			This is actually 3 operations:
				1. read counter value
				2. increment counter
				3. store value
			--> this can lead to errors if multiple threads are doing this at once unless we use a mutex
		*/
		counterMutex.Lock()
		counter++
		counterMutex.Unlock()
		wg.Done()
	}()
}

func updateCounterMap(id int) {
	wg.Add(1)
	go func() {
		mapCounterMutex.Lock()
		mapCounter++
		m[id] = id
		mapCounterMutex.Unlock()
		wg.Done()
	}()
}
