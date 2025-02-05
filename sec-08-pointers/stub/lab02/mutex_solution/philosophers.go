package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Philosopher: 0 1 2 3 4
fork: 		0 1 2 3 4 0
*/
const numPhilophers = 9

var (
	waiter sync.Mutex
	forks  [numPhilophers]sync.Mutex
	wg     sync.WaitGroup
	start  = time.Now()
)

func main() {

	for i := 0; i < numPhilophers; i++ {
		left := i
		right := i + 1
		if i == numPhilophers-1 {
			right = 0
		}
		eat(i, left, right)
	}

	wg.Wait() // wait forever

}

func eat(id, left, right int) {
	wg.Add(1)
	go func() {
		for {

			// call waiter
			waiter.Lock()

			// fmt.Printf("Philosopher %v got the waiter's attention\n", id)

			// pick up forks
			gotLeft := forks[left].TryLock()
			gotRight := forks[right].TryLock()

			if gotLeft && gotRight {
				fmt.Printf("%v - Philosopher %v is enjoying their meal and has dismissed the waiter\n", time.Since(start), id)
				waiter.Unlock()

				time.Sleep(time.Duration((rand.Intn(5))+1) * time.Second)

				fmt.Printf("%v - Philosopher %v has finished eating and is putting down their forks\n", time.Since(start), id)
				forks[left].Unlock()
				forks[right].Unlock()
			} else {
				// fmt.Printf("Philosopher %v couldn't eat! Dismissing the waiter and putting down forks\n", id)
				waiter.Unlock()
				if gotLeft {
					forks[left].Unlock()
				}
				if gotRight {
					forks[right].Unlock()
				}
			}

			// think before calling waiter again
			time.Sleep(time.Duration((rand.Int()%1000)+1) * time.Millisecond)

		}
	}()

}
