package main

import "fmt"

type (
	ID     string
	Person struct {
		name string
	}
)

func main() {
	for i := 0; i < 10; i++ {
		myPrintln(getValue())
	}

	// Pitfall - can't use .(type) outside of switch
	// var foo = getValue()
	// fooType := foo.(type)
}

func myPrintln(x any) {
	switch x.(type) { // checks type of x
	case bool:
		fmt.Printf("[bool] value: %v\n", x)
	case int:
		fmt.Printf("[int] value: %v\n", x)
	case float64:
		fmt.Printf("[float64] value: %v\n", x)
	case complex128:
		fmt.Printf("[complex128] value: %v\n", x)
	case *Person:
		fmt.Printf("[*Person] value: %v\n", x)
	default:
		fmt.Printf("<unknown> - value: %v, type: %T\n", x, x)
	}
}

// create value of random type
func getValue() any {
	ch := make(chan any, 1)
	select {
	case ch <- 2010:
	case ch <- 9.15:
	case ch <- "Hello world":
	case ch <- &Person{"John Doe"}:
	case ch <- 7 + 7i:
	case ch <- ID("19520925"):
	case ch <- true:
	}

	return <-ch
}
