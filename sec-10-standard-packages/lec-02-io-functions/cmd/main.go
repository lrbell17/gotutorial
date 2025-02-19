package main

import (
	"fmt"
	"io"
	"os"

	// "github.com/lrbell17/gotutorial/sec-10-standard-packages/lec-02-io-reader-and-writer/io"

	"github.com/lrbell17/gotutorial/sec-10-standard-packages/lec-02-io-functions/ms"
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

	WriteString(rw, "Hello, world")
	WriteString(rw, ". ")
	WriteString(rw, "It is a wonderful day.")

	Copy(os.Stdout, rw) // copy buffer to stdout, calling the Read method in the background

	fmt.Println() // should print "Hello, world. It is a wonderful day."
}

// WriteString converts a string to []bytes and then writes it to an io.Writer
func WriteString(w io.Writer, s string) (n int, err error) {
	if w == nil {
		return 0, io.ErrShortWrite
	}

	return w.Write([]byte(s))
}

func Copy(dst io.Writer, src io.Reader) (n int64, err error) {

	// check for nil src / dest
	if dst == nil || src == nil {
		return 0, io.ErrNoProgress
	}

	// create a buffer
	p := make([]byte, 256)

	for err == nil {
		c, readErr := src.Read(p) // read from src

		if c > 0 {
			c, err = dst.Write(p[:c]) // write what we have read to dst
			n += int64(c)
		}

		if readErr == io.EOF {
			return
		}

		if readErr != nil {
			return n, readErr
		}
	}
	return
}
