package temp

import (
	"fmt"

	"github.com/lrbell17/gotutorial/shared/input"

	"math/rand"
)

const (
	dataSize = 5000
	minTemp  = -20
	maxTemp  = 120
)

var tempData []int

func init() {
	var tempDataArray [dataSize]int
	for i := range tempDataArray {
		tempDataArray[i] = input.GetRandInt(minTemp, maxTemp)
	}
	tempData = tempDataArray[:]
}

func bubbleSort() {
	for i := 0; i < dataSize-1; i++ {
		for j := i + 1; j < dataSize; j++ {
			if tempData[j] < tempData[i] {
				tempData[i], tempData[j] = tempData[j], tempData[i]
			}
		}
	}
}

func getAvg() int {
	var sum int
	for _, v := range tempData {
		sum += v
	}
	return sum / dataSize
}

/*
pivot:
  - correct position is final in sorted array
  - items to the left are smaller
  - items to the right are larger
*/
func quickSort(arr []int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)

		// recursively sort before and after pivot
		quickSort(arr, low, pivotIndex-1)
		quickSort(arr, pivotIndex+1, high)
	}
}

// rearrange slices into elements less than and greater than the pivot
func partition(arr []int, low, high int) int {

	/*
		Chose a pivot and move it to the end
		Note: there are many different ways to choose a pivot (e. high, low, median of 3)
	*/
	pivotIndex := rand.Intn(high-low+1) + low // get random pivot index
	arr[pivotIndex], arr[high] = arr[high], arr[pivotIndex]
	pivot := arr[high] // pivot = last element

	i := low // index of smaller element

	/*
		iterate through range of partition
		if the current value is smaller than the pivot, swap values and increment index of smaller element
	*/
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	// swap back pivot element
	arr[i], arr[high] = arr[high], arr[i]

	return i // return new pivot index
}

func Print() {
	// bubbleSort()
	quickSort(tempData, 0, dataSize-1)
	fmt.Printf("Sorted data: %v\n", tempData)
	fmt.Printf("min: %v, max: %v, avg: %v\n", tempData[0], tempData[dataSize-1], getAvg())
}
