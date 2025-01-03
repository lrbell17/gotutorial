package main

import "fmt"

/*
Eveything is a number (video, doc, pictures, strings, etc)
  - a number is a collection of bits
  - a bit can store a value of 0 or 1 (like a switch)
  - just 1 bit acn store 2 distinct states or values
  - multiple bits can store more values (2^n)

Grouping bits
  - 4 bits = nibble
  - 8 bits = byte
  - 16 bits
  - 32 bits
  - 64 bits - most modern computers use this many bits for addressing memory
  - 128 bits

What is data?
  - data is another name for value

What is data type?
  - defines the interpretation and representation of the value (e.g. integer, float, string, etc)
  - defines the valid operations (e.g. addition, concatenation, etc)

Signed vs unsigned
  - signed numbers can be positive or negative
  - unsigned numbers are always positive

Basic data types
  - Boolean
  - true/false
  - String
  - Numeric
    --> Integer: 8-bit, 16-bit, 32-bit, 64-bit
    --> Floating point: 32-bit, 64-bit
    --> Complex numbers: 64-bit, 128-bit

Numeric data types:
  - Integer:
    --> unint8 (0-255), uint16 (0-65535), uint32 (0-4294967295), uint64 (0-18446744073709551615)
    --> int8 (-128-127), int16 (-32768-32767), int32 (-2147483648-2147483647), int64 (-9223372036854775808-9223372036854775807)
  - Floating point:
    --> float32, float64
  - Complex numbers:
    --> complex64 (float32 real and imaginary parts), complex128 (float64 real and imaginary parts)
  - Byte: alias for uint8
  - Rune: alias for int32
*/

// declaration at package level
const (
	pi                      = 3.141592
	ageLimit          uint8 = 18 // override default type
	defaultCountry          = "USA"
	isDebugingEnabled       = false
	_favChar                = 'A'
)

var (
	sot  = 9
	dir  = "20 deg south"
	name = "VLA"
)

func main() {
	fmt.Println(pi)

	//  default types: float64, int, string, bool, int32
	fmt.Printf("pi: %v, %T\n", pi, pi)
	fmt.Printf("ageLimit: %v, %T\n", ageLimit, ageLimit)
	fmt.Printf("defaultCountry: %v, %T\n", defaultCountry, defaultCountry)
	fmt.Printf("isDebugingEnabled: %v, %T\n", isDebugingEnabled, isDebugingEnabled)
	fmt.Printf("_favChar: %v, %T\n", _favChar, _favChar)

	// variables (can be changed)
	var a bool
	fmt.Println(a) // default value is false
	a = true
	fmt.Println(a)

	/*
	 Shorthand declaration for variables
	 - can only be used inside a function
	 - cannot specify type
	*/
	i := 7
	i = i + 9
	fmt.Println(i)
	i += 10
	fmt.Println(i)

	// declare multiple variables
	var t, u uint8 = 20, 30
	fmt.Println(t, u)
	x, y, z := 40, 5.0, "this is a string"
	fmt.Println(x, y, z)

	// type conversion
	g := 365.4
	d := g + 9
	fmt.Println(d)
	var e = int(g + 9)
	fmt.Printf("e: %v, %T\n", e, e)

	// fmt.Println(int(pi)) // doesn't work
	pi2 := pi
	fmt.Println(int(pi2)) // this works
	const pi3 = 3.0
	fmt.Println(int(pi3)) // this works (only bc the decimal is 0)

	// Go allows you to write arbitrary large numbers, we just can't store them
	const huge = 9999999999999999999999999999999999999999999999999999999999999999999999999999
	//var z uint64 = huge // this doesn't work
	var huge2 = huge / 100000000000000000000000000000000000000000000000000000000000 // this works
	fmt.Println(huge2)

	// run another function
	foo()
}

func foo() {
	fmt.Println(pi)
}
