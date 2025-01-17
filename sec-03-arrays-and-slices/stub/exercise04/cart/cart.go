package cart

import (
	"fmt"
	"math/rand"
)

const numCarts = 20

var carts []Currency

func init() {
	var cartArray [numCarts]Currency
	for i := range cartArray {
		cartArray[i] = getTotal()
	}
	carts = cartArray[:]
}

func quickSort(arr []Currency, low, high int) {
	if low < high {
		partitionIndex := patition(arr, low, high)

		quickSort(arr, partitionIndex+1, high)
		quickSort(arr, low, partitionIndex-1)
	}
}

func patition(arr []Currency, low, high int) int {

	pivotIndex := rand.Intn(high-low+1) + low
	arr[pivotIndex], arr[high] = arr[high], arr[pivotIndex]
	pivot := arr[high]

	i := low

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]

	return i

}

func processCart(arr []Currency) {
	for _, v := range arr {
		discount := v * 0.05
		tax := v * 0.08
		total := v - discount + tax
		fmt.Printf("Original price: %v, disccount: %v, tax: %v, total: %v\n",
			v, discount, tax, total)
	}
}

func Print() {
	quickSort(carts, 0, numCarts-1)
	fmt.Println(carts)
	processCart(carts)
}
