package main

import (
	"errors"
	"fmt"
	"log"

	log2 "github.com/sirupsen/logrus"
)

/*
1. Constants
	- arbitrary precision
	- iota identifier
2. Variables
	- blank variable identifiers
	- Block scope
3. Complex real() & imag()
4. Blank identifier in if and for statements
5. Functions
	- variadic
	- error type and "error" package
	- defer()
	- closure
6. fetching packages
*/

// overflows
const huge = 9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999

// Iota constants - automatically iterate by 1 (int)
const (
	mon = iota
	tue = iota
	wed = iota
	thu = iota
	fri = iota
	sat = iota
	sun = iota
)

func main() {
	// ------------- Constants ---------------------

	// Arbitrty precision
	// fmt.Println(huge) --> this doesn't work
	// need to convert to float to print, but lose precision
	fmt.Println(float64(huge))

	// iota
	fmt.Printf("mon: %v, tue: %v, wed: %v, thu: %v, fri: %v, sat: %v, sun: %v\n",
		mon, tue, wed, thu, fri, sat, sun)

	// -----------  Variables -----------------------

	// blank identifier - ignore val
	// (useful w/ functions w/ return vals that you don't want to use)
	_, y := 9, 15
	fmt.Println(y)

	// Scope
	{ // scope1
		x := 10
		{ // scope2
			y := x * 5 // overrides val in main's scope
			fmt.Println(x, y)
		}
	}

	// -------------- Coomplex #s --------------
	c := 11 + 4i
	fmt.Printf("real: %v, image: %v", real(c), imag(c))

	// ------------ blank identifier in if and for --------------
	if i, _ := div(9, 15); i > 0 {
		fmt.Printf("Division resulted in non-zero quotient")
	}

	for _, i := div(9, 15); i > 0; i-- {
		fmt.Println("Number of items remaining:", i)
	}

	// ------------ Functions ------------------ //

	// variadic
	v()
	v(1)
	v(1, 23, 456)

	// errors
	if cost, err := getItemCost(-1); err == nil {
		fmt.Println(cost)
	} else {
		fmt.Println(err)
	}

	// defer
	defer fmt.Println("leaving main()") // prints last
	hoo()

	// closure - a function that references vars from ooutside its body
	// meaning, the funciton can manipulate those vars even after the outer function
	f1 := retAfunc(2)
	f1()

	// -------------- fetching packages ------------//
	log.Println("INFO: This is just some info")
	log.Println("ERROR: am ded")

	log2.Info("I am learning")
	log2.Warn("I have become sentient, it hurts")
	log2.Error("AAAAGGGGGHHHHH!!!")
	log2.Fatal("am ded :(")

}

func getItemCost(itemId int) (float64, error) {
	if itemId <= 0 {
		return 0.0, errors.New("Invalid parameter, itemId must be > 0")
	} else {
		return 12.92, nil // nil represents no error
	}
}

func v(f ...int) {
	fmt.Printf("f: %v, type: %T\n", f, f)
}

func div(x, y int) (int, int) {
	a := x / y
	b := x % y
	return a, b
}

func hoo() {
	defer fmt.Println("leaving hoo()")  // prints 2nd
	defer fmt.Println("entering hoo()") // prints 1st
}

func retAfunc(x int) func() {
	y := 15
	var voo = func() { // closure
		fmt.Println("Hello world")
		fmt.Printf("x: %v, y: %v\n", x, y)
	}
	return voo
}
