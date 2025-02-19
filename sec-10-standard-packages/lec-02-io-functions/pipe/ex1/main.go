package main

import (
	"fmt"
	"sync"
)

var data = []string{"hello world!", "this is very nice", "how cool!"}

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)
	producer(&wg, ch)
	consumer(&wg, ch)
	wg.Wait()
}

func consumer(wg *sync.WaitGroup, in <-chan string) {
	wg.Add(1)

	go func() {
		defer wg.Done()
		for s := range in {
			fmt.Println(s)
		}
	}()

}

// producer uses write-only channel
func producer(wg *sync.WaitGroup, out chan<- string) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, s := range data { // write data to channel
			out <- s
		}
		close(out)
	}()
}
