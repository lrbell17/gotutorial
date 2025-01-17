package main

import (
	"fmt"

	"github.com/lrbell17/gotutorial/shared/input"
)

const dataSize = 100

type dataSet [dataSize]int

var data dataSet

func main() {
	fmt.Println(data)
	fmt.Printf("Min value: %v\n", min(data))

	sortedData := sort(data)
	fmt.Printf("Sorted data: %v\n ", sortedData)

}

func init() {
	for i := 0; i < len(data); i++ {
		data[i] = input.GetRandInt(-10, 20)
	}
}

func min(arr dataSet) int {
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		val := arr[i]
		if val < min {
			min = val
		}
	}
	return min
}

/*
Bubble sort
- compare 2 nums i,j to see what smallest, then swap
- iterate j by 1
- after all js, iterate i
*/
func sort(arr dataSet) dataSet {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
