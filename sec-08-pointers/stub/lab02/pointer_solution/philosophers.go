package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numPhilophers = 5
	dinnerTime    = 10 * time.Second
)

var (
	wg     sync.WaitGroup
	waiter = &Waiter{forks: make(chan *Fork, numPhilophers)}
)

type (
	Waiter struct {
		forks chan *Fork
	}
	Fork struct{}
)

func main() {

	for i := 0; i < numPhilophers; i++ {
		waiter.forks <- &Fork{}
		philosopher(i)
	}

	wg.Wait() // wait forever

}

func philosopher(id int) {
	wg.Add(1)

	go func() {
		defer wg.Done()
		end := time.After(dinnerTime)
		for {
			select {
			case <-end:
				return
			default:
				if rand.Intn(2)%2 == 0 {
					eat(id)
				} else {
					think(id)
				}
			}
		}

	}()
}

func think(id int) {
	fmt.Printf("Philosopher %v is thinking...\n", id)
	time.Sleep(time.Duration((rand.Int()%1000)+1) * time.Millisecond)
}

func eat(id int) {
	fmt.Printf("Philosopher %v wants to eat...\n", id)

	fork1, fork2 := waiter.getForks() // get forks from waiter

	fmt.Printf("Philosopher %v is enjoing their food!\n", id)
	time.Sleep(time.Duration((rand.Int()%1000)+1) * time.Millisecond)

	waiter.returnForks(fork1, fork2) // return forks to waiter
}

func (waiter *Waiter) getForks() (fork1, fork2 *Fork) {
	for {
		select {
		case fork1 = <-waiter.forks: // try to get fork1
			select {
			case fork2 = <-waiter.forks: // try to get 2nd fork
				return
			default:
				waiter.forks <- fork1 // return fork1 if can't get fork2
			}
		}
		time.Sleep(time.Duration((rand.Int()%1000)+1) * time.Millisecond) // wait random time before trying again
	}
}

func (waiter *Waiter) returnForks(fork1, fork2 *Fork) {
	waiter.forks <- fork1
	waiter.forks <- fork2
}
