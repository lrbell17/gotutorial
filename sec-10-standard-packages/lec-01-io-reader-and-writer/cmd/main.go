package main

import (
	"fmt"
	"io"
	"os"

	// "github.com/lrbell17/gotutorial/sec-10-standard-packages/lec-01-io-reader-and-writer/io"

	"github.com/lrbell17/gotutorial/sec-10-standard-packages/lec-01-io-reader-and-writer/ms"
)

type (
	// ReadWriter combines the interfaces io.Reader and io.Writer
	ReadWriter interface {
		io.Reader
		io.Writer
	}
)

func main() {
	var m, _ = ms.NewMemStore(100)
	test(m)
}

func test(rw ReadWriter) {
	// var totalBytes int
	// n, _ := rw.Write([]byte("Hello, world"))
	// totalBytes += n

	// n, _ = rw.Write([]byte(". "))
	// totalBytes += n

	// n, _ = rw.Write([]byte("It is a wonderful day."))
	// totalBytes += n

	// b := make([]byte, 1)
	// _, err := rw.Read(b)
	// for err == nil {
	// 	fmt.Print(string(b))
	// 	_, err = rw.Read(b)
	// }

	// io.WriteString calls the Write method in the background
	io.WriteString(rw, "Hello, world")
	io.WriteString(rw, ". ")
	io.WriteString(rw, "It is a wonderful day.")

	io.Copy(os.Stdout, rw) // copy buffer to stdout, calling the Read method in the background

	fmt.Println() // should print "Hello, world. It is a wonderful day."
}
