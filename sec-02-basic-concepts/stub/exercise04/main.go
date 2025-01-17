/*
	TODO 1 - Write a complete Go program which declares the

following identifiers as either const or var:

++ NOTE ++: Depending on the context, some of these might be
const or var. So there isn't a single correct answer.

	pi - Used instead of the value 3.14
	secondsInHour - Used instead of the value 60
	hoursInDay - Used instead of the value 24
	presenterName - name of current presenter which may change as records are processed
	favoriteLanguage - holds your favoritate programming or written language
	itemCount - represents the number of items in a shopping cart
	totalPrice - price of all items in a shopping cart
	isLoggingEnabled
*/
package main

import "fmt"

func main() {
	const pi = 3.14
	const secondsInHour = 60
	const hoursInDay = 24
	var presenterName string
	var favoriteLanguage string
	var itemCount int
	var totalPrice float64
	var isLoggingEnabled bool

	fmt.Println("pi:", pi)
	fmt.Println("secondsInHour:", secondsInHour)
	fmt.Println("hoursInDay:", hoursInDay)
	fmt.Println("presenterName:", presenterName)
	fmt.Println("favoriteLanguage:", favoriteLanguage)
	fmt.Println("itemCount:", itemCount)
	fmt.Println("totalPrice:", totalPrice)
	fmt.Println("isLoggingEnabled:", isLoggingEnabled)
}
