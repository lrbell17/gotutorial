package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const numWorkers = 100000

var wg sync.WaitGroup

/*
	Go can run many goroutines on a single OS thread
*/

func main() {

	for i := range numWorkers {
		producer(i)
	}
	wg.Wait()

}

func producer(id int) {
	wg.Add(1)
	go func() {
		fmt.Printf("Producer #%v is done\n", id)
		time.Sleep(time.Duration(rand.Int()%5000) * time.Nanosecond)
		wg.Done()
	}()
}
