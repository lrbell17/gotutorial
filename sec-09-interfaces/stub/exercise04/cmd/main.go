package main

import (
	"fmt"

	"github.com/lrbell17/gotutorial/sec-09-interfaces/stub/exercise04/wc"
)

type (
	// Writer interface declares the Write method for sending data to a value
	Writer interface {
		// Write takes a buffer b of type []bytes and returns the number of bytes accepted.
		// If n < len(b), then Write will return an appropriate error explaining why.
		Write(b []byte) (n int, err error)
	}
)

func main() {
	wc1 := new(wc.WriteCounter)
	storeData(wc1, data)
	fmt.Println(wc1)
}

// TODO 4 - implement the function storeData(), which takes a Writer and []string.
// storeData writes each string in the slice to the Writer
func storeData(w Writer, data []string) {

	for _, v := range data {
		w.Write([]byte(v))
	}
}
