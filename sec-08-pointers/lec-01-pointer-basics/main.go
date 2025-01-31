package main

import "fmt"

/*
Variables in memory:
- a unit of memory is a byte (8 bits per byte)

0:  xxxxxxxx00
10: xxxx000000
20: 0000000x00
30: 0000000000

consider variables:
  - name: price, type: float64, # bytes: 8, address in memory: 0 (start idx --> can infer end idx from start idx and type)
  - name: count, type: int32, # bytes: 4, address in memory: 10
  - name: isOn, type: bool, # bytes: 1, address in memory: 27

A pointer is the starting memory location of some value of a certain type --> pointer = address + type
  - 0, float64 --> *float64
  - 10, int32 --> *int32
  - 27, bool --> *bool

Formal definition: A pointer is an appropriate type that can hold the numeric value representing the location of a value and the type in memory
  - e.g. a pointer of type int can hold the number value representing the memory locaction of an int value
  - NOTE: we can only get the address of a non-const & non-literal value stored in memorys

we can also have a pointer of a pointer and so on ...
--> count int32, pCount *int32, ppCount **int32 ...
*/
func main() {

	// declaring a pointer
	var pCount *int
	fmt.Printf("pCount: %v, pCount: %p\n", pCount, pCount) // %p prints out hexidecimal represention of pointer

	// more examples
	var pBool *bool
	var pString *string
	var pInt64 *int64
	var pMap *map[string]int
	var pChan *chan int

	fmt.Printf("pBool: %v, pString: %v, pInt64: %v, pMap: %v, pChan: %v\n",
		pBool, pString, pInt64, pMap, pChan)
	fmt.Printf("pBool: %p, pString: %p, pInt64: %p, pMap: %p, pChan: %p\n",
		pBool, pString, pInt64, pMap, pChan)

	// Declaring pointers to user-defined types
	type SSN string
	var pSsn *SSN
	fmt.Printf("pSsn: %v, pSsn: %p\n", pSsn, pSsn)

	type Person struct {
		fname string
		lname string
		age   uint8
	}
	var pPerson *Person
	fmt.Printf("pPerson: %v, pPerson: %p\n", pPerson, pPerson)

	// --> all uninitialized pointers have value nil and memory address 0

	// A function name is pointer to the function
	var pFunc func(int)
	fmt.Printf("pFunc: %v, pFunc address: %p\n", pFunc, pFunc)

	//pFunc(7) --> get an error for nil pointer dereference
	// pFunc = &goo --> get an error bc we can't take address of goo
	pFunc = goo
	fmt.Printf("pFunc: %v, pFunc address: %p\n", pFunc, pFunc) //  the value matches the address
	pFunc(7)

	// nil is a speccial value without a type (until it needs a type)
	// assigning nil means that you are no longer using it (GC can clean it up)
	pBool = nil
	pString = nil
	pInt64 = nil
	pMap = nil
	pChan = nil

	fmt.Printf("pBool: %v, pString: %v, pInt64: %v, pMap: %v, pChan: %v\n",
		pBool, pString, pInt64, pMap, pChan)
	fmt.Printf("pBool: %p, pString: %p, pInt64: %p, pMap: %p, pChan: %p\n",
		pBool, pString, pInt64, pMap, pChan)

	// other things that are not pointers can still have nil value
	var s []int
	var m map[int]*Person
	var c chan SSN
	var f func(int, string, *Person) (int, complex64)
	fmt.Printf("s: %v, m: %v, c: %v, f: %v\n", s, m, c, f)
	fmt.Printf("s: %p, m: %p, c: %p, f: %p\n", s, m, c, f)
	s = nil
	m = nil
	c = nil
	f = nil
	fmt.Printf("s: %v, m: %v, c: %v, f: %p\n", s, m, c, f)
	fmt.Printf("s: %p, m: %p, c: %p, f: %p\n", s, m, c, f)

}

func goo(x int) {
	fmt.Printf("I am in goo: %v\n", x)
}
