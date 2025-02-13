package wc

import (
	"errors"
	"fmt"
)

// TODO 1 - declare type Writecounter
// WriteCounter keeps track of the following, number calls to Write() and the number of bytes written
type WriteCounter struct {
	byteCount  int
	writeCount int
}

// TODO 2 - implement the method Write(b []byte) (n int, err error)
// Write updates the value each time this method is called and the number of bytes written
func (wc *WriteCounter) Write(b []byte) (n int, err error) {
	if wc == nil {
		return 0, errors.New("invalid input - nil object")
	}

	n = len(b)
	wc.byteCount += n
	wc.writeCount++

	return
}

// TODO 3 - imeplement the method String() string
// String return an informative message about how the Write() was used
func (wc WriteCounter) String() string {
	return fmt.Sprintf("There were %v write operations totaling %v bytes", wc.writeCount, wc.byteCount)
}
