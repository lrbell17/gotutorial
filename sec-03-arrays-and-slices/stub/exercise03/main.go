package main

import (
	"fmt"

	"github.com/lrbell17/gotutorial/shared/input"
)

const (
	size   = 10
	minVal = -10
	maxVal = 10
)

type currency float64

var nums [size]currency

func (c currency) String() string {
	return fmt.Sprintf("$%.2f", c)
}

func main() {
	slice := nums[:]
	minPrice, maxPrice := slice[0], slice[0]
	var sum currency

	for _, v := range slice {
		sum += v
		if v < minPrice {
			minPrice = v
		}
		if v > maxPrice {
			maxPrice = v
		}
	}

	avg := sum / currency(len(slice))

	fmt.Println(slice)
	fmt.Printf("min: %v, max: %v, total: %v, avg: %v\n", minPrice, maxPrice, sum, avg)

}

func init() {
	for i := range nums {
		nums[i] = currency(input.GetRandFloat(minVal, maxVal))
	}
}
