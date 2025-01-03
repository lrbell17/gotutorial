package main

import "fmt"

func main() {
	fmt.Println("Hello cruel world!")

	/*
		A rune is a 32 bit number that represents a Unicode code point.
		A code point is how unicode characters are encoded.
		A rune is bit enough to represent any printable character, even in
		languages likke Chinese with 1000s of characters
	*/
	fmt.Println('H')      // returns a number! 72
	fmt.Println('H' + 20) // works with numeric operations

	fmt.Printf("19990 = %c\n", 19990) // prints the character represented by 19990
	fmt.Printf("19990 = %d\n", 19990) // print as decimal
	fmt.Printf("19990 = %X\n", 19990) // print as hexidecimal

	// https://asciitable.com

	// esacpe characters
	fmt.Println("hello", 10, "world")
	fmt.Println('\n')

	fmt.Printf("hello\nworld\n")
	fmt.Println('\\', "\\")
	fmt.Println('"', "\"")

	// Raw stings
	fmt.Println(`This is a raw string that are interpreted literally
which can span multiple lines
NOTE: escape characters are not \n\n interpreted`)

	// String concatenation
	fmt.Println("Hello" +
		" " +
		"World")
}
