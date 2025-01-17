package main

import "fmt"

func main() {

	// CECS = create, expand, copy, shrink

	/*
		Creating slicecs at runtime
		we always need an array before we have a slice
		--> can we have an anonymoous array we ccreate the slice from?
	*/

	// creating slice from array
	var s0 []int
	fmt.Printf("s0: %v, len: %v, type: %T\n", s0, len(s0), s0)

	var arr [10]int
	s0 = arr[:]
	fmt.Printf("s0: %v, len: %v, type: %T\n", s0, len(s0), s0)

	// make slice from function - not ideal bc function can't take size variable
	s1 := makeSlice10Int()
	s1[0] = -1
	fmt.Printf("s1: %v, len: %v, type: %T\n", s1, len(s1), s1)

	s2 := makeSlice10Int() //creates a whole separate slice
	s2[0] = 99
	fmt.Printf("s2: %v, len: %v, type: %T\n", s2, len(s2), s2)

	// Go has a build in make function for this
	s3 := make([]int, 5)
	s4 := make([]string, 7)
	fmt.Printf("s3: %v, len: %v, type: %T\n", s3, len(s3), s3)
	fmt.Printf("s4: %v, len: %v, type: %T\n", s4, len(s4), s4)

	/*
		Expanding slices
	*/
	var s5 []string
	fmt.Printf("s5: %v, len: %v, type: %T\n", s5, len(s5), s5)

	s5 = append(s5, "New York")                           //  need too assign to variable
	s5 = append(s5, "Chicago", "San Francisco", "Berlin") // Variadic!! can append multiple values
	fmt.Printf("s5: %v, len: %v, type: %T\n", s5, len(s5), s5)

	/*
		 	Slice shallow copy - editing slice will also edit original
			s1  := arr[:]
			s2 := s1
	*/
	var s6 = s5
	s6[0] = "Boston"
	fmt.Printf("s5: %v, len: %v, type: %T\n", s5, len(s5), s5)
	fmt.Printf("s6: %v, len: %v, type: %T\n", s6, len(s6), s6)

	/*
		Slice deep copy
		s1 := arr[:]
	*/
	s7 := make([]string, 3) // need to create desrination slice first

	// if length of source is larger than dest, it will only copy the smaller #
	copy(s7, s6)                                               // s7 --> 3, s6 --> 4
	fmt.Printf("s7: %v, len: %v, type: %T\n", s7, len(s7), s7) // returns len 3

	// if length of source is smaller, it will only copy that many records
	s8 := make([]string, 7)
	copy(s8, s6)                                               // s8 --> 7, s6 --> 4
	fmt.Printf("s8: %v, len: %v, type: %T\n", s8, len(s8), s8) // returns len 7 (4 vals, 3 empty strings)

	var s9 []string
	copy(s8, s9)
	fmt.Printf("s9: %v, len: %v, type: %T\n", s9, len(s9), s9) // no place to copy

	// instead, we can make a deep copy by appending
	s9 = append(s9, s8...)
	fmt.Printf("s9: %v, len: %v, type: %T\n", s9, len(s9), s9)

	/*
		Shrinking slices
	*/

	// shallow copy shrink
	s10 := s9[1:3]
	s10[0] = "Trenton"
	fmt.Printf("s9: %v, len: %v, type: %T\n", s9, len(s9), s9)
	fmt.Printf("s10: %v, len: %v, type: %T\n", s10, len(s10), s10)

	// deep copy shrink
	var s11 []string
	s11 = append(s11, s9[1:3]...)
	s11[0] = "Santa Cruz"
	fmt.Printf("s10: %v, len: %v, type: %T\n", s10, len(s10), s10)
	fmt.Printf("s11: %v, len: %v, type: %T\n", s11, len(s11), s11)
}

func makeSlice10Int() []int {
	// var arr [size]int --> this is not allowed bc size needs to be constant
	var arr [10]int
	return arr[:]
}
