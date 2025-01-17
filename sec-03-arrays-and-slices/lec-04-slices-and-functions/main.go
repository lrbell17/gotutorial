package main

import (
	"fmt"

	"github.com/lrbell17/gotutorial/shared/input"
)

const dataSize = 100

type dataSet []int

var data [dataSize]int

func main() {
	slice := data[:]

	fmt.Println(slice)
	fmt.Printf("Min value: %v\n", min(slice))

	sort(slice)
	fmt.Printf("Sorted data: %v\n ", data)

	// variadic function
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 5))

	// using the expand operator, we can pass a slice into the variadicc function
	// only works with slices, not arrays!
	fmt.Println(sum(data[:]...))

	// can also use the expand opeator to specify the length of an array
	arr2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(sum(arr2[:]...))

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
// sorting the slice will change the original array!!!
// advantage: we don't need to ccopy the whole array - good for time and memory
func sort(arr dataSet) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func sum(v ...int) (s int) {

	fmt.Printf("sum().v type is a slice: %T\n", v)
	for _, v := range v {
		s += v
	}
	return s
}
