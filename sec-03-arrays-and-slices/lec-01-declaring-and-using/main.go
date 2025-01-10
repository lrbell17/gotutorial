package main

import "fmt"

type Currency float64

func (c Currency) String() string {
	s := fmt.Sprintf("$%.2f", float64(c))
	return s
}

func main() {

	// before arrays
	var price0 Currency = 9.23
	var price1 Currency = 0.59
	var price2 Currency = 12.84
	var price3 Currency = 5.95
	var price4 Currency = 16.37
	var price5 Currency = 1.35

	fmt.Println(price0, price1, price2, price3, price4, price5)
	totalPrice := price0 + price1 + price2 + price3 + price4 + price5
	avgPrice := totalPrice / 6
	fmt.Printf("Total price: %v, Avg price: %v\n", totalPrice, avgPrice)

	// declaring an array
	var prices [6]Currency // length = 6
	fmt.Println(prices)    // all values default to 0.00

	// add elements (starts at 0, no negative elements)
	prices[0] = price0
	prices[1] = price1
	prices[2] = price2
	prices[3] = price3
	prices[4] = price4
	prices[5] = price5
	// prices[6] = 1.00 --> compilation error: index out of bounds

	fmt.Println(prices)

	// Resetting the varible later on does not affect the value stored in the array
	price0 = 99.99
	fmt.Println(prices)

	// Accecss values from array:
	fmt.Printf("Value: %v, Type: %T\n", prices[0], prices[0])

	// Iterating
	totalPrice = 0
	for i := 0; i < len(prices); i++ {
		totalPrice += prices[i]
	}
	avgPrice = totalPrice / Currency(len(prices))
	fmt.Printf("Total price: %v, Avg price: %v\n", totalPrice, avgPrice)

	// iterating using range operator
	for i, v := range prices { // can return index and value
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}

	// if we want, we can ignore the index with blank operator"_"
	totalPrice = 0
	for _, v := range prices { // can return index and value
		totalPrice += v
	}

	// Another way to intialize
	var prices2 = [10]Currency{ // don't need to fill all 10
		price0,
		price1,
		price2,
		price3,
		price4,
		price5,
	}
	fmt.Println(prices2)
}
