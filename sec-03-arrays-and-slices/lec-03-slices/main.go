package main

import "fmt"

type currency float64

func main() {
	// a slice is a subset or window of an array
	// --> array, offset, length

	var a [10]int
	fmt.Printf("a: %v, type: %T\n", a, a)

	var cart [10]currency
	fmt.Printf("a: %v, type: %T\n", cart, cart)

	// declaring a slice
	var s0 []int
	fmt.Printf("s0: %v, type: %T\n", s0, s0)

	var nums = [10]int{12, 1, 5, 99, 43, 65, 76, 87, 98, 9}
	fmt.Printf("nums: %v, len: %v, type: %T\n", nums, len(nums), nums)

	// [:] to grab all records
	// same as [0:len(nums)]
	s1 := nums[:]
	fmt.Printf("nums[:]: %v, len: %v, type: %T\n", s1, len(s1), s1)

	// get first 5 elements
	// same as nums[0:5]
	s2 := nums[:5]
	fmt.Printf("nums[:5]: %v, len: %v, type: %T\n", s2, len(s2), s2)

	// all elements after index 2
	// same as nums[2:len(nums)]
	s3 := nums[2:]
	fmt.Printf("nums[2:]: %v, len: %v, type: %T\n", s3, len(s3), s3)

	// upper and lower bound
	s4 := nums[3:7]
	fmt.Printf("nums[3:7]: %v, len: %v, type: %T\n", s4, len(s4), s4)

	/*
		Summary of indexing [x:y]
		x is inclusive
		y is not inclusice
		y - x = len
	*/

	// modifying the slices also modifies the underlying array and other slices!!
	for i := range s1 {
		s1[i] = i
	}
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("nums:", nums) // same result

}
