package main

import "fmt"

func stats(
	getNumOfItems func() int,
	getItemCost func(int) float64) {

	var totalCost, avgCost, maxCost, minCost float64
	minCost = getItemCost(1)
	var itemCount = getNumOfItems()

	for x := 1; x <= itemCount; x++ {
		c := getItemCost(x)
		totalCost += c

		if c > maxCost {
			maxCost = c
		}
		if c < minCost {
			minCost = c
		}
	}
	avgCost = totalCost / float64(itemCount)

	fmt.Printf("Number of items: %v\n", itemCount)
	fmt.Printf("Total price: $%.2f\n", totalCost)
	fmt.Printf("Average cost $%.2f\n", avgCost)
	fmt.Printf("Min cost: $%.2f\n", minCost)
	fmt.Printf("Max cost: $%.2f\n", maxCost)
}
