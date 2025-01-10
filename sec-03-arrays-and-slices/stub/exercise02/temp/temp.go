package temp

import (
	"fmt"

	"github.com/lrbell17/gotutorial/shared/input"
)

const (
	dataSize = 100
	minVal   = -20
	maxVal   = 120
)

type tempDataSet [dataSize]int

var tempData tempDataSet

func init() {
	for i := 0; i < dataSize; i++ {
		tempData[i] = input.GetRandInt(minVal, maxVal)
	}
}

func sort(data tempDataSet) tempDataSet {
	for i := 0; i < dataSize-1; i++ {
		for j := i + 1; j < dataSize; j++ {
			if data[j] < data[i] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	return data
}

func avg(tempData tempDataSet) int {

	var sum int
	for _, v := range tempData {
		sum += v
	}
	return sum / dataSize
}

func Print() {
	fmt.Printf("Data: %v\n", tempData)
	sortedTempData := sort(tempData)
	fmt.Printf("Sorted data: %v\n", sortedTempData)
	fmt.Printf("Min: %v, Max: %v, Avg: %v\n", sortedTempData[0], sortedTempData[dataSize-1], avg(tempData))
}
