package main

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

/*
We want to be able to send/recieve to/from multiple channels from one goroutine

The "select" keyword picks one of the channels that is ready
*/
func main() {

	// inserting delays
	fmt.Printf("Message 1 at: %v\n", time.Now())
	time.Sleep(1 * time.Second) // pauses the current goroutine
	fmt.Printf("Message 2 at: %v\n", time.Now())

	// getting notification from a channel
	fmt.Printf("Message 1 at: %v\n", time.Now())
	alarm := notifyAfter(1 * time.Second)
	<-alarm
	fmt.Printf("Message 2 at: %v\n", time.Now())

	// Go's time package has a similar built-in function
	fmt.Printf("Message 1 at: %v\n", time.Now())
	<-time.After(1 * time.Second)
	fmt.Printf("Message 2 at: %v\n", time.Now())

	// the select statement (looks same as switch)
	var ch1 chan int
	select {
	case <-ch1:
		log.Info("Got data from ch1")
	default:
		log.Warn("No data from ch1") // having a default ensures there is no block
	}

	// selecting data from multiple channels
	// the select statement (looks same as switch)
	var ch2, ch3 chan string
	select {
	case <-ch2:
		log.Info("Got data *FROM* ch2")
	case ch3 <- "hello":
		log.Info("Sent data *TO* ch3")
	default:
		log.Warn("No comms for ch2 or ch3")
	}

	// Generating random bit stream
	fmt.Print("Randon bit stream: ")
	total := 100000
	bits := randomBitsGen(total)

	// validating bit stream
	m := make(map[int8]int)
	for v := range bits {
		fmt.Print(v)
		m[v]++
	}
	fmt.Println()
	for k, v := range m {
		f := float32(v) / float32(total) * 100
		fmt.Printf("%v occcurred %.2f%% of the time\n", k, f)
	}

	// timing out after wating on a channel
	d := producer()
	cousumer(d)

}

func producer() (out chan string) {
	out = make(chan string)
	go func() {
		for i := 0; i < 1000; i++ {
			out <- fmt.Sprintf("Producer message %v", i+1)
		}
		close(out)
	}()
	return
}

func cousumer(in chan string) {
	for {
		timeout := time.After(1 * time.Millisecond)

		select {
		case m, ok := <-in:
			if !ok {
				fmt.Println("No more data coming from 'in'")
				return
			}
			fmt.Printf("Cousumer got: %v\n", m)
		case <-timeout:
			fmt.Println("Timeout waiting for data")
			return
		}
	}
}

func notifyAfter(delay time.Duration) (out chan time.Time) {
	out = make(chan time.Time)
	go func() {
		time.Sleep(delay)
		out <- time.Now() // add current time to channel after sleeps
		close(out)
	}()
	return
}

func randomBitsGen(l int) (out chan int8) {

	// serial implementation
	// out = make(chan int8, l) // make channel of length l
	// defer close(out) // defer close to end of func
	// for {
	// 	select {
	// 	case out <- 0:
	// 	case out <- 1:
	// 	default:
	// 		return // return once channel is full
	// 	}
	// }

	// concurrent implementation
	out = make(chan int8) // unbuffered channel
	go func() {
		for i := 0; i < l; i++ {
			// select performs uniform distribution of cases
			select {
			case out <- 0:
			case out <- 1:
			case out <- 1: // can skew results by making more cases
			}
		}
		close(out)
	}()
	return
}

// Incorrect way to insert delay - gorouotine is still working!
// func sleep(delay time.Duration) {
// 	end := time.Now().Add(delay)
// 	for time.Now().Before(end) {
// 		log.Debug("Just waiting...")
// 	}
// }
