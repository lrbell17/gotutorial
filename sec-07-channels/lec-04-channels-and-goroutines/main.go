package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var (
	wgProducers sync.WaitGroup
	wgConsumers sync.WaitGroup
)

/*
sender --------- (channel) ---------> receiver
- sender blocked when receiver dies
- receiver blocked when sender dies (and if it's a buffered channel, once receiver drains all the data)
*/
func main() {

	// buffered channel
	// ch := make(chan string, 5)
	// go producer(1, ch)
	// consumer(ch)

	// unbuffered channel
	ch2 := make(chan string)

	// start consumers
	consumer(1, ch2)
	consumer(2, ch2)

	// start producers
	producer(1, ch2)
	producer(2, ch2)
	producer(3, ch2)

	wgProducers.Wait() //  wait for all producers
	close(ch2)         // close channel
	wgConsumers.Wait() // wait for consumers
}

func producer(id int, out chan string) {

	wgProducers.Add(1)

	// publish messages to channel for 1s
	go func() {
		i := 1
		end := time.Now().Add(1000 * time.Millisecond)

		for time.Now().Before(end) {
			out <- fmt.Sprintf("Producer: %v, item: %v", id, i)
			i++
		}
		wgProducers.Done()
	}()
}

func consumer(id int, in chan string) {
	wgConsumers.Add(1)

	go func() {
		db := make(map[string]int)
		var fields []string
		for v := range in {
			fields = strings.Split(v, ",")
			db[fields[0]]++
		}

		for k, v := range db {
			fmt.Printf("Consumer %v - processed %v items for %v\n", id, v, k)
		}

		wgConsumers.Done()
	}()
}
