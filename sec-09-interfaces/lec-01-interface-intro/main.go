package main

import "fmt"

/*
An interface defines a set of operations/methods allowed
--> like a contract, whatever implements the interface needs to define the methods
--> a type can implement or provide operations for several interfaces

Interfacce vs Struct
  - struct defines the memory layout of a type
    --> deals with fields and types
    --> can't put methods inside of structs
  - interface defines which operations are available
    --> deals with methods
    --> can't put fields inside an interface

*/

type (
	// commonly, one method interface ends with "er" and then the method is the same w/o the "er"
	Reader interface {
		Read(buf []byte) (n int, err error)
	}
	Writer interface {
		Write(buf []byte) (n int, err error)
	}

	// good practice to limit the # of methods per interface
	Shape interface {
		Area()
	}
	Circle interface {
		Radius()
		Shape // Shape interface is embedded
	}
	Rect interface {
		Width() float64
		Length() float64
		Shape
	}
)

func main() {

	/*
		Empty interface
	*/
	type Empty interface{}

	var e Empty

	// both value AND type default to nil
	fmt.Printf("e's value: %v, type: %T\n", e, e)

	// can store any type in the Empty interface
	e = 7
	fmt.Printf("e's value: %v, type: %T\n", e, e) // type = int

	e = "Hello world"
	fmt.Printf("e's value: %v, type: %T\n", e, e) // type = string

	f := 3.14
	e = &f
	fmt.Printf("e's value: %v, type: %T\n", e, e) // type = *float64

	var dataSource Reader
	var printer Writer

	fmt.Printf("dataSource value: %v, type: %T\n", dataSource, dataSource)
	fmt.Printf("printer value: %v, type: %T\n", printer, printer)

}
