package main

// lukebell@managers-MacBook-Pro db % go doc
// package db // import "github.com/lrbell17/gotutorial/sec-02-basic-concepts/lec-11-packages-p1/db"
// go install
import (
	"fmt"

	"github.com/lrbell17/gotutorial/sec-02-basic-concepts/lec-11-packages-p1/db"
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
	stats(db.GetNumOfItemsDb1, db.GetItemCostDb1)
	stats(db.GetNumOfItemsDb2, db.GetItemCostDb2)
}
