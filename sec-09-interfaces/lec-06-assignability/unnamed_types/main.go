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

func main() {

	// assignment between named and unnamed types
	var p0 Person
	var p1 struct {
		name string
		age  Age
	}

	// an unnamed type is assignable if it is EXACTLY like the named type
	p0 = p1
	p1 = p0

	// unnamed type structs are not assignable if the order of fields are different
	var p2 struct {
		age  Age
		name string
	}
	// p2 = p0 --> error
	// p0 = p2 --> error
	fmt.Println(p2)

	// structs with different # of fields
	var p3 struct {
		name string
		age  Age
		ssn  ID
	}

	// p0 = p3 --> error
	// p3 = p0 --> error
	fmt.Println(p3)

	// struct types are different if field names are different
	var p4 struct {
		Name string
		Age  Age
	}
	// p0 = p4 --> error
	// p4 = p0 --> error
	fmt.Println(p4)

}
