package main

import "fmt"

func main() {
	// 6, true, 2.14 are literal values
	// 3 + 1, 2+2 == 4 is an expression

	// boolean
	fmt.Println("Booleans: ")
	fmt.Println(true)
	fmt.Println(false)
	fmt.Println("true && false = ", true && false)
	fmt.Println("true || false = ", true || false)
	fmt.Println("!true = ", !true)
	fmt.Println("5 > 10 = ", 5 > 10)
	fmt.Println("abc == xyz = ", "abc" == "xyz")

	// Numeric
	fmt.Println("\nNumerics: ")
	fmt.Println(6)
	fmt.Println("3 + 1 = ", 3+1)
	fmt.Println("-20: ", -20) // negative number
	fmt.Println("-7 + 10 =", -7+10)
	fmt.Println("-7 - -10 =", -7 - -10) // double neg evaluates to positive
	fmt.Println("-7-(-10) =", -7-(-10))
	fmt.Println("-7 + -10 =", -7+-10)
	// Integer division
	fmt.Println("9 / 15 =", 9/15) // evaluates to 0
	// Modulus
	fmt.Println("9 % 15 =", 9%15) // evaluates to 9

	// Examples of floating point numbers
	fmt.Println("\nFloating Point Numbers: ")
	fmt.Println("9.0 / 15 =", 9.0/15) // as long as one of the number is float, the result is float
	fmt.Println("9 / 15.0 = ", 9/15.0)
	fmt.Println("Pi ~=", 3.141592)

	// Go supports complex numbers
	fmt.Println("\nComplex numbers: ")
	fmt.Println("1 + 2i = ", 1+2i)
	fmt.Println("25i = ", 25i)
	fmt.Println("real(11+4i) =", real(11+4i)) // evaluates to 11
	fmt.Println("imag(11+4i) =", imag(11+4i)) // evaluates to 4
}
