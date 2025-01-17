package main

import "fmt"

// define currency type (needs to be based on another type (e.g. primive))
type (
	currency float64
	greeter  func() string // can have function types
)

func formatCurrency(c currency) (s string) {
	s = fmt.Sprintf("$%.2f", float64(c)) // returns string instead of printing
	return
}

// method - Go is an object-oriented lanuage
// c is a reciever of type currency
func (c currency) formatCurrency() (s string) {
	s = fmt.Sprintf("$%.2f", float64(c)) // returns string instead of printing
	return
}

// if we override the String method (same as toString method)
func (c currency) String() (s string) {
	s = fmt.Sprintf("$%.2f", float64(c)) // returns string instead of printing
	return
}

func main() {
	var price1 currency = 99.99
	var price2 currency = 32.49
	total := price1 + price2

	// Need to format into 2 decimal places
	fmt.Printf("Total price: $%.2f\n", total)
	avgPrice := avg(price1, price2)
	fmt.Printf("Average price: $%.2f, type: %T\n", avgPrice, avgPrice)

	// use format functions
	fmt.Printf("Total price (formatted): %v\n", formatCurrency(total))
	fmt.Printf("Avgg price (formatted): %v\n", formatCurrency(avgPrice))

	// receiver method:
	fmt.Printf("Total price (method): %v\n", total.formatCurrency())
	fmt.Printf("Avg price (method): %v\n", avgPrice.formatCurrency())

	// print w/ String method:
	fmt.Printf("Total price (String): %v\n", total)
	fmt.Printf("Avg price (String): %v\n", avgPrice)

	// ------------------------------
	fmt.Println(hi())
	fmt.Println(goodMorning())

	say(hi, "world")
	say(goodMorning, "Luke")
}

func avg(v1 currency, v2 currency) (avg currency) {
	avg = (v1 + v2) / 2.0
	return
}

func hi() string {
	return "hi"
}

func goodMorning() string {
	return "Good Morning"
}

func say(greet greeter, s string) {
	fmt.Printf("%v, %v\n", greet(), s)
}
