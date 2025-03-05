package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type (
	Message struct {
		Source  string
		Payload []int
	}
)

func main() {
	wg := new(sync.WaitGroup)
	ch := producer(wg, "Producer 1")
	consumer(wg, ch)
	wg.Wait()
}

func consumer(wg *sync.WaitGroup, in <-chan Message) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		timeout := time.After(5 * time.Second)

		for {
			select {
			case msg := <-in:
				processMessage(msg)
				timeout = time.After(5 * time.Second) // reset timer after we receive work
			case t := <-timeout:
				log.Warnf("Waited %v to send message, shutting down producer", t)
				return
			}
		}
	}()
	return
}

func producer(wg *sync.WaitGroup, src string) (out <-chan Message) {
	wg.Add(1)
	ch := make(chan Message)
	out = ch
	go func() {
		defer wg.Done()
		timeout := time.After(5 * time.Second)
		for {
			select {
			case ch <- newMessage(src):
				time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond) // random delay up to 2sec
				timeout = time.After(5 * time.Second)                         // reset timer after sending

			case t := <-timeout:
				log.Warnf("Waited %v to send message, shutting down producer", t)
				close(ch)
				return
			}
		}
	}()
	return

}

// create random payload
func newMessage(src string) Message {
	count := rand.Intn(10) + 1 // 1 to 10
	data := make([]int, count)
	for i := 0; i < count; i++ {
		data[i] = rand.Intn(100)
	}
	return Message{Source: src, Payload: data}
}

// calculate stats for message
func processMessage(msg Message) {
	total := 0

	for v := range msg.Payload {
		total += v
	}
	fmt.Printf("%v send:\n", msg.Source)
	l := len(msg.Payload)
	fmt.Printf("\tcount: %v, total: %v, avg: %v\n", l, total, float64(total)/float64(l))
}
