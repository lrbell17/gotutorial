package stats

import "fmt"

type Currency float64

func (c Currency) String() string {
	s := fmt.Sprintf("$%.2f", float64(c))
	return s
}

func GetTotal(prices []Currency) (totalPrice Currency) {

	for _, v := range prices {
		totalPrice += v
	}
	return
}

func GetMax(prices []Currency) (maxPrice Currency) {
	maxPrice = prices[0]
	for _, v := range prices {
		if v > maxPrice {
			maxPrice = v
		}
	}
	return
}

func GetMin(prices []Currency) (minPrice Currency) {
	minPrice = prices[0]
	for _, v := range prices {
		if v < minPrice {
			minPrice = v
		}
	}
	return
}

func GetAvg(prices []Currency) (avgPrice Currency) {
	avgPrice = GetTotal(prices) / Currency(len(prices))
	return
}
