package main

import (
	"fmt"
)

const isLoggingEnabled = true

func main() {

	// simple if statement
	if isLoggingEnabled {
		fmt.Println("yes, log this info")
	}

	// if statement with arithmetic
	x := 5
	if x > 10 {
		fmt.Printf("x > 10: %v\n", true)
	} else {
		fmt.Printf("x > 10: %v\n", false)
	}

	// short var declaration in 'if' statement
	if x := 4; x < 5 {
		fmt.Println("x is local to the 'if' block and is different from 'x' in main. x:", x)
	}

	fmt.Println("x:", x)

	// else if
	if x < 5 {
		fmt.Println("x < 5")
	} else if x > 10 {
		fmt.Println("x > 10")
	} else {
		fmt.Println("all tests fail")
	}

	// Strings
	if 'A' < 'B' {
		fmt.Println("'A' < 'B'")
	}

	if "aa" < "A" {
		fmt.Println("\"aa\" < \"A\"")
	}

	if today := "Wednesday"; today == "Wednesday" {
		fmt.Println("It is Wednesday my dudes, AAAAAAAAAAAA!!!")
	}

}
