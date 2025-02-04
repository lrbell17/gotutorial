package main

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"strings"
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
	usedBytes     int
}

func main() {

	aDoc, _ := loadDoc("dummy document")
	printDoc(aDoc)

	fmt.Printf("Size of aDoc object: %v\n", unsafe.Sizeof(aDoc))
	fmt.Printf("Word count: %v\n", countWords(aDoc))
	capitalizeHeading(aDoc, HEADING1)
	listHeadings(aDoc, HEADING1)

	printDoc(aDoc)

	// example passing a pointer of a pointer - not really useful
	var bDoc *Document
	loadDocDoublePointer(&bDoc, "Document 2")
	printDoc(bDoc)
}

func countWords(doc *Document) (words int) {
	if doc == nil {
		return
	}
	fmt.Printf("Size of object passed through countWords: %v\n", unsafe.Sizeof(doc))

	// ...

	return rand.Int() % 200000
}

func capitalizeHeading(doc *Document, level string) {
	if doc == nil {
		return
	}
	fmt.Printf("Size of object passed through capitalizeHeading: %v\n", unsafe.Sizeof(doc))
	doc.ModifiedAt = time.Now()

	s := string(doc.Data[:doc.usedBytes])
	s = strings.ToTitle(s)
	var buf bytes.Buffer
	fmt.Fprintln(&buf, s)
	doc.usedBytes, _ = buf.Read(doc.Data[:100])

}

func listHeadings(doc *Document, level string) {
	fmt.Printf("Size of object passed through listHeadings: %v\n", unsafe.Sizeof(doc))
	// ...

}

func loadDoc(src string) (doc *Document, err error) {
	doc = new(Document)
	if doc == nil {
		return nil, errors.New("Cannot allocate new document")
	}

	doc.Name = src
	doc.Author = "John Smith"
	doc.CreatedAt = time.Now()
	doc.ModifiedAt = time.Now()
	var buf bytes.Buffer
	fmt.Fprintln(&buf, "chapter 1")
	doc.usedBytes, _ = buf.Read(doc.Data[:100])
	return
}

// load doc using a pointer of a pointer - don't need to return the doc
func loadDocDoublePointer(ppDoc **Document, src string) (err error) {

	if ppDoc == nil {
		return errors.New("Invalid parameter value for ppDoc")
	}
	doc := new(Document)
	*ppDoc = doc
	if doc == nil {
		return errors.New("Cannot allocate new document")
	}

	doc.Name = src
	doc.Author = "John Smith"
	doc.CreatedAt = time.Now()
	doc.ModifiedAt = time.Now()
	var buf bytes.Buffer
	fmt.Fprintln(&buf, "chapter 1")
	doc.usedBytes, _ = buf.Read(doc.Data[:100])
	return
}

func printDoc(doc *Document) {
	if doc == nil {
		return
	}
	fmt.Printf("\nDoc name: %v\n", doc.Name)
	fmt.Printf("Doc created at: %v\n", doc.CreatedAt)
	fmt.Printf("Doc modified at: %v\n", doc.ModifiedAt)
	fmt.Printf("%-6v %v\n", 1, string(doc.Data[:doc.usedBytes]))
}
