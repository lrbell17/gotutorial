package main

import (
	"fmt"
)

func main() {

	// s
	a, b := 1, 2
	fmt.Printf("a: %v, b: %v\n", a, b)
	a, b = swap(a, b)
	fmt.Printf("a: %v, b: %v\n", a, b)

	a, b = 3, 4
	fmt.Printf("a: %v, b: %v\n", a, b)
	a, b = swap2(a, b)
	fmt.Printf("a: %v, b: %v\n", a, b)

	// variatic
	sum()
	sum(1, 2)
	sum(1, 2, 3, 4, 5)

	formatter("format 1", 1.0, 1.1, 1.2)

	// anonymous function is called the same way as a regular function
	foo()
	fmt.Printf("foo's type: %T\n", foo) // returns func()
	hi("Luke", 3)
	fmt.Printf("hi's type: %T\n", hi) // returns func(string, int)

	// can define anonymous func as variable
	var hi2 func(string, int)
	fmt.Printf("hi2's value: %v\n", hi2) // returns nil
	hi2 = hi
	fmt.Printf("hi2's value: %v\n", hi2) // retuens hex value
	hi2("world", 2)

	// anonymous functions make is easy to reuse code
	stats(getNumOfItemsDb1, getItemCostDb1)
	stats(getNumOfItemsDb2, getItemCostDb2)
}

// parameters both int
func swap(s, t int) (int, int) {
	return t, s
}

// can also specify which values we cant to return
func swap2(s, t int) (x int, y int) { // function signature
	// function body starts here
	x = t
	y = s
	// return 10, 25 --> something like this is still valid
	return
}

// variatic functions: can include many parameters of a given type
func sum(v ...int) {
	fmt.Printf("v values are: %v\n", v)
	fmt.Printf("v's type is: %T\n", v) // []int
}

// variatic parameter must be the last parameter
func formatter(s string, f ...float64) {
	// do stuff
}

// anonymous function - define funciton as variable
var foo = func() {
	fmt.Println("Im am the foo func")
}

var hi = func(s string, c int) {
	for i := 0; i < c; i++ {
		fmt.Printf("Hi, %v\n", s)
	}
}

/// reuse code from prev lecctures

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

// ------- DON"T LOOK BELOW HERE
var db1 = map[int]float64{
	1: 340.77, 2: 883.66, 6: 1411.14, 9: 1639.36, 10: 1481.56,
	3: 1084.18, 4: 577.13, 5: 796.16, 7: 224.36, 8: 1138.59,
}
var db2 = map[int]float64{
	1: 1098.12, 2: 519.92, 3: 628.44, 4: 126.35, 5: 697.27,
	6: 852.28, 7: 1576.11,
}

// getNumOfItemsDb1 returns the totalNumber of items in database 1
func getNumOfItemsDb1() int {
	return len(db1)
}

/*
	getItemCostDb1 returns the cost of an item identified by 'itemID',

where 'itemID' is 1 to getNumofItemsDb1()
*/
func getItemCostDb1(itemID int) (c float64) {
	c = db1[itemID]
	return
}

// getNumOfItemsDb2 returns the totalNumber of items in database 2
func getNumOfItemsDb2() int {
	return len(db2)
}

/*
	getItemCostDb2 returns the cost of an item identified by 'itemID',

where 'itemID' is 1 to getNumofItemsDb2()
*/
func getItemCostDb2(itemID int) (c float64) {
	c = db2[itemID]
	return
}
