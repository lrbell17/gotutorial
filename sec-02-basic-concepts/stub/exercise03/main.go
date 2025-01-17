/*
TODO 1 - Write a complete Go program which prints the
'Featured articles' on the Go Language website: https://golang.org
TIP : Copy the text and print it using a raw string.
*/
package main

import "fmt"

func main() {
	fmt.Println(`Talks
A Video Tour of Go¶
Three things that make Go fast, fun, and productive: interfaces, reflection, and concurrency. Builds a toy web crawler to demonstrate these.

Code that grows with grace¶
One of Go's key design goals is code adaptability; that it should be easy to take a simple design and build upon it in a clean and natural way. In this talk Andrew Gerrand describes a simple "chat roulette" server that matches pairs of incoming TCP connections, and then use Go's concurrency mechanisms, interfaces, and standard library to extend it with a web interface and other features. While the function of the program changes dramatically, Go's flexibility preserves the original design as it grows.

Go Concurrency Patterns¶
Concurrency is the key to designing high performance network services. Go's concurrency primitives (goroutines and channels) provide a simple and efficient means of expressing concurrent execution. In this talk we see how tricky concurrency problems can be solved gracefully with simple Go code.

Advanced Go Concurrency Patterns¶
This talk expands on the Go Concurrency Patterns talk to dive deeper into Go's concurrency primitives.

More¶
See the Go Talks site and wiki page for more Go talks.`)
}
