package main

import "fmt"

func main() {

	// switch statement
	// case is just like an "if-else", when one condition is satisfied, it exits the switch block
	// need to have "fallthrough" in order to continue
	// just like "if", x is only avilable in the switch scope
	switch x := 3; {
	case x < 5:
		fmt.Println("x is less than 5")
		fallthrough
	case x < 4:
		fmt.Println("x is less than 4")
	case x > 10:
		fmt.Println("x is greater than 10")
	case x == 5:
		fmt.Println("x is equal to 5")
	default:
		fmt.Println("default")
	}

	// shorthand is useful for function calls
	//switch day := getDay(); {

	// Strings
	switch day := "Saturday"; day { //specify variable to match against
	case "Monday":
		fmt.Println("Ugh")
	case "Tuesday", "Wednesday": // works as an OR operator
		fmt.Printf("Today is %v, time to werk it\n", day)
	case "Thursday":
		fmt.Println("Uno mas")
	case "Friday":
		fmt.Println("TGIF")
	case "Saturday", "Sunday":
		fmt.Printf("%v is a day for relaxation\n", day)
	default:
		fmt.Printf("What day is this? %v\n", day)
	}
}
