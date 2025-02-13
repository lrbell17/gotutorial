package main

import "fmt"

/*
Type Assertion: mechanism to reveal dynamic type
Type Switching: using type as switch case condition

Static vs dynamic type:
var foo interface{}
foo = 3.14

What is foo's type?
--> static type is interface{}
--> dynamic type is float64
*/
func main() {
	var foo interface{}
	foo = 1971.07

	var f float64

	// f = float64(foo) --> error, needs type assertiono
	f = foo.(float64)
	fmt.Printf("f's value: %v, type: %T\n", f, f)

	foo = 77
	// f = foo.(float64) --> panic: interface conversion: interface {} is int, not float64
	// fmt.Printf("f's value: %v, type: %T\n", f, f)

	/*
		Checking if the type assertion is correct
	*/
	foo = getValue()

	// 2nd return type for checking if the type assertion is correct
	intVal, ok := foo.(int)
	fmt.Printf("int? %v, intVal's value: %v\n", ok, intVal) // defaults to 0

	floatVal, ok := foo.(float64)
	fmt.Printf("float? %v, floatVal's value: %v\n", ok, floatVal) // defaults to 0

	strVal, ok := foo.(string)
	fmt.Printf("string? %v, strVal's value: %v\n", ok, strVal) // defaults to ""

	pVal, ok := foo.(*Person)
	fmt.Printf("*Person? %v, pVal's value: %v\n", ok, pVal) // defaults to nil

}

type Person struct {
	name string
}

// create value of random type
func getValue() any {
	ch := make(chan any, 1)
	select {
	case ch <- 2010:
	case ch <- 9.15:
	case ch <- "Hello world":
	case ch <- &Person{"John Doe"}:
	}

	return <-ch
}
