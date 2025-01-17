package main

import "fmt"

/*
	A map is an unordered group of elements of one type
	Indexed by a set of unique keys (can be another type)

	The value of an unititialized map is nil
*/

func main() {

	// initialize slicce for comparison
	var s []int
	if s != nil {
		fmt.Printf("s: %v, len: %v, type: %T\n", s, len(s), s)
	}

	// Declaring a map
	var m map[int]string
	fmt.Printf("m: %v, len: %v, type: %T\n", m, len(m), m)

	// m[98] = "A+" --> doesn't work since m is not initialized (hence is nil)

	// create map using make
	m = make(map[int]string)
	fmt.Printf("m: %v, len: %v, type: %T\n", m, len(m), m)

	// Add values
	m[98] = "A+"
	m[95] = "A"
	m[-1] = "Invalid"

	fmt.Printf("m: %v, len: %v, type: %T\n", m, len(m), m)
	fmt.Printf("m[95]: %v\n", m[95])

	// key not found - no error - just returns default value for type (empty string in this case)
	fmt.Printf("m[70]: %v, type: %T\n", m[70], m[70])

	m1 := make(map[string][]string)
	m1["Smith"] = make([]string, 4)
	m1["Smith"][0] = "Mary"
	m1["Smith"][1] = "John"
	m1["Smith"][2] = "Pete"
	m1["Smith"][3] = "Anne"

	fmt.Printf("Smith family: %v\n", m1["Smith"])

	// nil is returned if the key does not exists
	fmt.Printf("Jones family: %v, type: %T\n", m1["Jones"], m1["Jones"]) // returns
	if m1["Jones"] == nil {
		fmt.Println("Confirmed - Jones is nill")
	}

	// however, we can also have a key with nil
	m1["Bell"] = []string{}
	fmt.Printf("Bell family: %v, type: %T\n", m1["Bell"], m1["Bell"]) // returns
	if m1["Jones"] == nil {
		fmt.Println("Confirmed - Bell is nill")
	}

	// how to conclusively determine if key is present
	if _, ok := m1["Jones"]; !ok {
		fmt.Println("Jones is not in the map")
	}

	// length of map
	fmt.Printf("length of m1: %v\n", len(m1))

	// iterating throuh map --> order is not guaranteed!!!
	for k, v := range m {
		fmt.Printf("m[%v] = %v\n", k, v)
	}

	// delete elements
	delete(m, -1)
	fmt.Println(m)

	// when define an array from another, we get 2 distinct copies
	var a [10]int
	var a1 = a
	a[0] = -1
	a1[0] = 2
	fmt.Printf("a: %v, a1: %v\n", a, a1)

	// however, when we do the same with a map, both maps are affected
	var m2 = m
	m2[70] = "C"
	fmt.Printf("m: %v, m2: %v\n", m, m2)

	// same behavior as slices
	var s1 = make([]int, 10)
	var s2 = s1
	s1[0] = -1
	s2[0] = 2
	fmt.Printf("s1: %v, s2: %v\n", s1, s2)
}
