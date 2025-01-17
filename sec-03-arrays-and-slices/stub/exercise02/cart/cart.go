package cart

import "fmt"

const (
	totalItems = 20
)

var shoppingCarts [totalItems]Currency

func init() {
	for i := range shoppingCarts {
		shoppingCarts[i] = getTotal()
	}
}

func adjustPrice(carts []Currency) {
	for _, v := range carts {
		discount := v * 0.05
		tax := v * 0.08
		total := v - discount + tax
		fmt.Printf("Original price: %v, discount: %v, tax: %v, total: %v\n", v, discount, tax, total)
	}
}

func Print() {
	adjustPrice(shoppingCarts[:])
}
