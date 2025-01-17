// Section 03 - Excercise 05

package main

import (
	"fmt"
	"log"
	"math/rand"
)

func main() {
	s0 := []string{"To", "be", "or", "not", "to", "be.", "That", "is", "the", "question."}
	s1 := make([]string, 10)
	myCopyStringSlice(s1, s0)
	fmt.Println(s1)

	testMyCopyStringSlice()
	var s []string
	s = myAppendStringToSlice(s, "I")
	s = myAppendStringToSlice(s, "love", "programming", "in", "Golang!")
	fmt.Println(s) // [I love programming in Golang!]

	arr := make([]int, 100)
	for i := range arr {
		arr[i] = int(rand.Int31n(100))
	}
	fmt.Println(arr)

	// bubbleSort(arr)
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

// TODO 1 - implemenat myCopyStringSlice()
func myCopyStringSlice(dest []string, src []string) {

	for i := 0; i < getMin(len(src), len(dest)); i++ {
		dest[i] = src[i]
	}
}

func getMin(val1, val2 int) int {
	if val1 > val2 {
		return val2
	} else {
		return val1
	}
}

// TODO 2 - Impelment your custom []stirng append function myAppendStringToSlice()
func myAppendStringToSlice(s []string, val ...string) []string {
	newLen := len(s) + len(val)
	newSlice := make([]string, newLen)

	myCopyStringSlice(newSlice, s)

	valIdx := 0
	for i := len(s); i < newLen; i++ {
		newSlice[i] = val[valIdx]
		valIdx++
	}

	return newSlice
}

// --- NO need to look below here
// TODO 1 - Implement your custom []string copy function myCopyStringSlice()
func testMyCopyStringSlice() {
	s0 := []string{"To", "be", "or", "not", "to", "be.", "That", "is", "the", "question."}
	var s1 []string
	myCopyStringSlice(s1, s0)
	s2 := make([]string, 4)
	expected := make([]string, 4)
	myCopyStringSlice(s2, s0[0:0])
	copy(expected, s0[0:0])
	for i, v := range s2 {
		if v != expected[i] {
			log.Fatalf("test in testMyCopyStringSlice() failed, expected: %v, got: %v", expected, s2)
		}
	}
	myCopyStringSlice(s2, s0[4:])
	copy(expected, s0[4:])
	for i, v := range s2 {
		if v != expected[i] {
			log.Fatalf("test in testMyCopyStringSlice() failed, expected: %v, got: %v", expected, s2)
		}
	}
}

func bubbleSort(arr []int) {

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func quickSort(arr []int, low, high int) {
	if low < high {
		partIdx := partition(arr, low, high)

		quickSort(arr, low, partIdx-1)
		quickSort(arr, partIdx+1, high)
	}
}

func partition(arr []int, low, high int) int {

	// get pivot
	// pivotIdx := high
	pivotIdx := low + rand.Intn(high-low+1)

	// swap high w/ pivot
	arr[high], arr[pivotIdx] = arr[pivotIdx], arr[high]
	pivot := arr[high]

	// lower bound
	i := low

	// iterate through whole partition
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	// swap back pivot
	arr[i], arr[high] = arr[high], arr[i]

	return i
}
