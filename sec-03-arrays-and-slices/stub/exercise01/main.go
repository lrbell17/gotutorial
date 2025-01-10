package main

import (
	"fmt"

	"github.com/lrbell17/gotutorial/sec-03-arrays-and-slices/stub/exercise01/stats"
	"github.com/lrbell17/gotutorial/shared/input"
)

const (
	arrLen = 20
	minVal = 0
	maxVal = 100
)

var prices [arrLen]stats.Currency

func main() {
	fmt.Println(prices)
	fmt.Printf("Total price: %v\n", stats.GetTotal(prices[:]))
	fmt.Printf("Min price: %v\n", stats.GetMin(prices[:]))
	fmt.Printf("Max price: %v\n", stats.GetMax(prices[:]))
	fmt.Printf("Avg price: %v\n", stats.GetAvg(prices[:]))
}

func init() {
	for i := 0; i < arrLen; i++ {
		prices[i] = stats.Currency(input.GetRandFloat(minVal, maxVal))
	}
}
