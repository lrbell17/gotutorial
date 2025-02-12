package wc

import (
	"errors"
	"fmt"
)

type (
	// WriteCounter keeps track of how many times the method Write() was called
	WriteCounter uint
	// Writer       interface {
	// 	Write([]byte)
	// }
)

// TODO 1 - implement the method Write(b []byte) (n int, err error) for WriteCounter
// Write updates the value each time this method is called
func (wc *WriteCounter) Write(b []byte) (n int, err error) {

	if wc == nil {
		return 0, errors.New("nil input")
	}

	*wc++
	return len(b), nil
}

// TODO 2 - imeplement the method String() string for WriteCounter
// String return an informative message about how many times the method Write() was called
func (wc WriteCounter) String() string {
	return fmt.Sprintf("There were %v write operations", int(wc))
}
