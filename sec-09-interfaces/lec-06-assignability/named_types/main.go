package main

import "fmt"

type (
	ID     string
	SKU    string
	Person struct {
		name string
		age  Age
	}
	Individual struct {
		name string
		age  Age
	}
	Age uint8
)

/*
 */
func main() {

	// assignable if types are exactly the same
	var i0 int16 = 1971
	var i1 int16
	i1 = i0
	fmt.Printf("i0: %v, type: %T, i1: %v, type: %T\n", i0, i0, i1, i1)

	// names types are ALWAYS different
	id0 := ID("71-11-04")
	sku := SKU("71-07-07")

	// id0 = sku --> error
	fmt.Printf("id0: %v, type: %T, sku: %v, type: %T\n", id0, id0, sku, sku)

	// structs are also named types - follow the same rule
	var jane Person
	var bob Individual

	// jane = bob --> error
	jane = Person(bob) // need to cast
	fmt.Printf("jane: %v, type: %T, bob: %v, type: %T\n", jane, jane, bob, bob)

}
