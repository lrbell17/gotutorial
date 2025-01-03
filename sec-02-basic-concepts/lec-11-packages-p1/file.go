package main

import "fmt"

// variatic functions: can include many parameters of a given type
func sum(v ...int) {
	fmt.Printf("v values are: %v\n", v)
	fmt.Printf("v's type is: %T\n", v) // []int
}

// variatic parameter must be the last parameter
func formatter(s string, f ...float64) {
	// do stuff
}

// anonymous function - define funciton as variable
var foo = func() {
	fmt.Println("Im am the foo func")
}

var hi = func(s string, c int) {
	for i := 0; i < c; i++ {
		fmt.Printf("Hi, %v\n", s)
	}
}
