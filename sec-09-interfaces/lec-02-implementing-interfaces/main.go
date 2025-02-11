package main

import "fmt"

type (
	Currency float64
	Stringer interface {
		String() string
	}
)

// A method is a function that takes a receiver
func (c Currency) String() string {
	return fmt.Sprintf("$%.2f", float64(c))
}

/*
Interface implementation is IMPLICIT
--> A type implements an interface when it provides definitions for ALL methods in the interface
--> there is no "implements" keyword
*/
func main() {
	var c Currency = 11.04
	fmt.Printf("c.String(): %v\n", c.String()) // explicit call to Currency's String method

	// Currency ALSO implements interface fmt.Stringer
	fmt.Printf("c: %v\n", c) // implicitly call Currency's String() method

	// Assigning values to interfaces - only works if the type implements all the methods in the interface!!
	var mainStringer Stringer = c
	fmt.Printf("mainStringer's value: %v, type: %T\n", mainStringer, mainStringer)

	var fmtStringer Stringer = c
	fmt.Printf("fmtStringer's value: %v, type: %T\n", fmtStringer, fmtStringer)

	// Also works for pointers
	var c1 = new(Currency)
	*c1 = 11.04
	fmt.Println(c.String())
	fmt.Println(c)

}
