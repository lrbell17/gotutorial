package main

import "fmt"

/*
A channel provides a method for concurrently executing functions
  - works by sending and recceiving values of a specified element type
  - value of unititialized channel is nil

FIFO: First in, first out

producer1()                				 consumer1()
producer2() --- data (via channel) ----> consumer2()
producer3()
*/
func main() {

	// declaring channels
	var ch chan int
	fmt.Printf("ch: %v, len: %v, cap: %v\n", ch, len(ch), cap(ch))

	// sending data to a channel
	// ch <- 4 --> get error that all goroutines are asleep - deadlock!

	// receiving data from a channel
	// <-ch --> get same error bc channel is nil

	/*
		Unbuffered channel (a channel without capacity)
	*/
	ch = make(chan int)
	fmt.Printf("ch: %v, len: %v, cap: %v\n", ch, len(ch), cap(ch))

	// sending data to an unbuffered channel also fails when there is no receiver
	// ch <- 2

	// receiving data to an unbuffered channel also fails when there is no sender
	// <-ch

	/*
		Buffered channel (a channel with capacity)
	*/
	ch = make(chan int, 2)
	ch <- 2
	ch <- 3
	// ch <- 4 --> fails bc we exceed channel capacity
	fmt.Printf("ch: %v, len: %v, cap: %v\n", ch, len(ch), cap(ch))

	v1 := <-ch
	v2 := <-ch
	// <-ch --> fails bc we read all the values already
	fmt.Printf("v1: %v, v2: %v, ch: %v, len: %v, cap: %v\n", v1, v2, ch, len(ch), cap(ch))

	/*
		Send-only and Receive-only channels
	*/
	chs := make(chan string, 5)

	// Send-only
	var out chan<- string = chs
	out <- "hello"
	out <- "world"

	// Receive-only
	var in <-chan string = chs // assign to same channel
	fmt.Println(<-in)
	fmt.Println(<-in)

}
