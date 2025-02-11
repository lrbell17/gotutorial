package main

import "fmt"

type (
	Currency float64
	Stringer interface {
		String() string
	}
)

// same method for pointers
func (c *Currency) String() string {
	return fmt.Sprintf("$%.2f", float64(*c))
}

/*
Interfaces and pointers
*/
func main() {
	var c = new(Currency)
	*c = 11.04
	fmt.Println(c.String())
	fmt.Println(c)

	// when we change the function to expect a pointer, we need to explicitly call the function
	var c1 Currency = 1.25
	fmt.Println(c1.String())
	fmt.Println(c1)

}
