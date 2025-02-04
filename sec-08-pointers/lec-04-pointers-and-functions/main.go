package main

import (
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

const (
	HEADING1    = "Heading-1"
	HEADING2    = "Heading-2"
	HEADING3    = "Heading-3"
	HEADING4    = "Heading-4"
	docDataSize = 100e6
)

type Document struct {
	Name          string
	Author        string
	CreatedAt     time.Time
	ModifiedAt    time.Time
	Data          [docDataSize]byte
	NewLineIndex  map[int]int
	PageLineIndex map[int]int
	headingIndex  map[string][]int
}

func main() {

	// passing pointers to functions
	count := 1737

	fmt.Printf("Count BEFORE incInt(): %v\n", count)
	incInt(count)
	fmt.Printf("Count AFTER incInt(): %v\n", count) // nothing happens bc only the copy of the count var is incremented

	// instead we can pass the pointer to a function
	incIntPointer(&count)
	fmt.Printf("Count AFTER incInt(): %v\n", count)

	// Document example - each time, the a cioy of the whole doc is passed into each function
	var aDoc Document
	fmt.Printf("Size of aDoc object: %v\n", unsafe.Sizeof(aDoc))
	fmt.Printf("Word count: %v\n", countWords(aDoc))
	capitalizeHeading(aDoc, HEADING1)
	listHeadings(aDoc, HEADING1)

}

func countWords(doc Document) (words int) {
	fmt.Printf("Size of object passed through countWords: %v\n", unsafe.Sizeof(doc))

	// ...

	return rand.Int() % 200000
}

func capitalizeHeading(doc Document, level string) {
	fmt.Printf("Size of object passed through capitalizeHeading: %v\n", unsafe.Sizeof(doc))
	doc.ModifiedAt = time.Now()
	// ...

}

func listHeadings(doc Document, level string) {
	fmt.Printf("Size of object passed through listHeadings: %v\n", unsafe.Sizeof(doc))
	// ...

}

func incInt(v int) {
	v++
}

func incIntPointer(v *int) {
	*v++
}
