package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Infinite loop
	// for true {
	// 	fmt.Println("PLEASE KILL ME")
	// }

	// Behaves like a while loop
	x := 1
	totalMessages := 5
	for x <= totalMessages {
		fmt.Printf("This is message %v of %v\n", x, totalMessages)
		x++
	}

	// Behaves like a for loop
	for x := 1; x <= totalMessages; x++ {
		fmt.Printf("This is message %v of %v\n", x, totalMessages)
	}

	// calculation of sample data
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

	fmt.Printf("Items over $1000.00: %v\n", getItemsMoreThan(1000))
}

func getItemsMoreThan(cutoffPrice float64) int {

	var countMoreThan int
	for x := 1; x <= getNumOfItems(); x++ {
		c := getItemCost(x)
		if c < cutoffPrice {
			continue
		}
		countMoreThan++
	}
	return countMoreThan
}

// ------- DON"T LOOK BELOW HERE
const dataPoints = 10

var db map[int]float64

// getNumOfItems returns the totalNumber of items in our database
func getNumOfItems() int {
	return dataPoints
}

// getItemCost returns the cost of an item identified by 'itemID', where 'itemID' is 1 to getNumofItems()
func getItemCost(itemID int) (c float64) {
	c = db[itemID]
	return
}

// init initializes our database of items and prices
func init() {
	db = make(map[int]float64)
	for i := 1; i <= dataPoints; i++ {
		rand.Seed(time.Now().Unix() + rand.Int63())
		// get a random float value
		v := rand.Uint32()
		v %= 200000                // limit range to no more than 200,000
		db[i] = float64(v) / 100.0 // results in 2 decimal places, price between $0 to $2,000.00
	}
}
