package main

import (
	"fmt"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	chCap = 10
)

var (
	s = rand.NewSource(time.Now().Unix())
	r = rand.New(s)
)

func main() {

	ch1 := make(chan int, chCap)
	if ch1 == nil {
		log.Error("Could not create channel")
	}

	// when we pass a channel into a function, it is only a shallow copy
	producer(ch1)
	consumer(ch1)

	// processor/mapper needs to return another channel
	ch2 := generator()
	ch2 = counter(ch2)
	ch2 = adder(ch2, 5)
	consumer(ch2)

}

func producer(nums chan int) {

	n := r.Int()%cap(nums) + 1 // between 1 and capacity
	log.Infof("Producing %v elements\n", n)

	for i := 0; i < n; i++ {
		nums <- r.Int() % 200 // between 0 and 199
	}
	close(nums)
}

func consumer(nums chan int) {
	for v := range nums {
		fmt.Printf("Consumer got %v\n", v)
	}
}

func generator() (out chan int) {
	log.Info("Generator of random ints")
	out = make(chan int, chCap)
	producer(out)
	return
}

// processor (a copy of the channel is returned)
func counter(in chan int) (out chan int) {
	out = make(chan int, len(in))
	count := 0
	for v := range in {
		out <- v
		count++
	}
	close(out)
	log.Infof("Counted %v elements\n", count)
	return
}

// mapper (a modified channel is returned)
func adder(in chan int, numToAdd int) (out chan int) {
	log.Infof("Adding %v to all elements\n", numToAdd)

	out = make(chan int, len(in))
	for v := range in {
		out <- v + numToAdd
	}
	close(out)
	return
}
