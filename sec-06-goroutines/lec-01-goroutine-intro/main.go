package main

import (
	"fmt"
	"time"
)

/*
	Sequential functions: f1 runs, then f2 runs
	|	f1 ------------> return
	P0
	|	f2						------------> return

	Concurrent functions: f1 and f1 run on same processor
	|	f1 ----->	  --->		---->
	P0
	|	f2		 ----> 	  ---->

	Parallelism: f1 and f1 run at the same time on multiple proccecssors
	P0	|	f1 -------------------->

	P1	|	f2 -------------------->

	Concurrency != Parallelism
		- Concurrency is a way of writing code
		- Parallelism is how the code executes
		- You get parallelism for free using concurrent programming when you have multiple cores
*/

func main() {

	// creating goroutine from named function
	go producer(1)
	go producer(2)

	// creating goroutine from anonymous function call
	go func(str string) {
		fmt.Printf("Foo() - Hello %v\n", str)
		for i := 0; i < 5; i++ {
			fmt.Printf("Foo() - message #%v\n", i)
		}
		producer(3)
	}("World") // call it immediately

	// give goroutine time to complete work (race condition where main() completes before goroutines execute)
	time.Sleep(1 * time.Millisecond)
}

func producer(id int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Producer #%v - message #%v\n", id, i)
	}
}
