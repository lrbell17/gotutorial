package main

import "fmt"

func main() {

	sc := make(chan string, 2)

	/*
		Sending to closed channel
	*/
	sc <- "hello"
	close(sc)

	fmt.Printf("sc: %v, len: %v, cap: %v\n", sc, len(sc), cap(sc))

	// sc <- "world" --> fails bc channel is closed

	/*
		Receiving from closed channel
	*/
	sc = make(chan string, 2)
	sc <- "hello"
	sc <- "world"

	s1 := <-sc
	fmt.Printf("s1: %v, sc: %v, len: %v, cap: %v\n", s1, sc, len(sc), cap(sc)) // s1 = hello

	close(sc) // close channel

	s2 := <-sc
	fmt.Printf("s2: %v, sc: %v, len: %v, cap: %v\n", s2, sc, len(sc), cap(sc)) // s2 = world
	s3 := <-sc
	fmt.Printf("s3: %v, sc: %v, len: %v, cap: %v\n", s3, sc, len(sc), cap(sc)) // s3 = nil

	/*
		checking for value sent before closure
	*/
	sc2 := make(chan string, 10)
	sc2 <- "hi"
	sc2 <- ""
	close(sc2)

	s, ok := <-sc2
	fmt.Printf("s: \"%v\", ok: %v\n", s, ok) // ok: true
	s, ok = <-sc2
	fmt.Printf("s: \"%v\", ok: %v\n", s, ok) // ok: true
	s, ok = <-sc2
	fmt.Printf("s: \"%v\", ok: %v\n", s, ok) // ok: false

	/*
		Incorrect ways of iterating over channels values
	*/
	fillCh(5, 3)
	// for i := 0; i < cap(ch); i++ { -> results in deadlock bc we try to read more values
	// for i := 0; i < len(ch); i++ { --> get incorrect # of records bc the length shrinks each iteration
	l := len(ch)
	for i := 0; i < l; i++ { // could be incorrect bc another goroutine could be updating the length
		fmt.Println(<-ch)
	}

	// range operator
	fillCh(5, 3)
	// for v := range ch { // get a deadlock bc range keeps trying to read from channel even after we consume all records
	// 	fmt.Println(v)
	// }

	/*
		Correct way of iterating through channel
	*/
	close(ch) // close channel first!
	for v := range ch {
		fmt.Println(v)
	}
}

var (
	ch chan int
)

func fillCh(c, l int) {
	ch = make(chan int, c)
	for i := 0; i < l; i++ {
		ch <- i
	}
}
