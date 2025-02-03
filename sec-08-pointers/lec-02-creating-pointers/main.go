package main

import "fmt"

func main() {

	// Review from last message --> creates a nil pointer
	var pCount *int
	fmt.Printf("pCount: %v, pCount: %p, type: %T\n", pCount, pCount, pCount)

	// The "address of" operator is "&"
	// Give a variable / value --> returns a pointer (address & type)
	var count = 1737
	pCount = &count
	fmt.Printf("pCount: %v, pCount: %p, type: %T\n", pCount, pCount, pCount)

	name := "John Doe"
	fmt.Printf("&name: %v, name: %v\n", &name, name)

	pCount = getAddressOfVar(count) // --> won't work since a COPY of the variable is passed in the function

	count2 := 1234
	pCount2 := &count2
	ppCount2 := &pCount2
	fmt.Printf("count value: %v, count address: %v, type: %T\n", count2, &count2, count2)
	fmt.Printf("pCount value: %v, pCount address: %v, type: %T\n", pCount2, &pCount2, pCount2)      // value matches address of count2
	fmt.Printf("ppCount value: %v, ppCount address: %v, type: %T\n", ppCount2, &ppCount2, ppCount2) // value matches address of pCount2

	// Use the new() funciton to get the address of memory allocated for the requested type
	// e.g. new(string) --> Address: 560(*string) --> pString: 560(*string)
	// same as s: "" string --> & --> Address: 560(*string) --> pStrin: 560(*string)

	pCount3 := new(int) // same as count3 := 0; pCount3 := *count
	fmt.Printf("pCount3 value: %v, pCount3 type: %T\n", pCount3, pCount3)

	pName := new(string) // same as name := ""; pName := *name
	fmt.Printf("pName value: %v, pName type: %T\n", pName, pName)

	pArr := new([10]int)
	fmt.Printf("pArr value: %v, pArr type: %T\n", pArr, pArr)

	pSlice := new([]float64)
	fmt.Printf("pSlice value: %v, pSlice type: %T\n", pSlice, pSlice)

	pMap := new(map[int64][]complex128)
	fmt.Printf("pMap value: %v, pMap type: %T\n", pMap, pMap)

	pCh := new(chan *chan string)
	fmt.Printf("pCh value: %v, pCh type: %T\n", pCh, pCh)

	// Using new() with custom types
	type ID string
	pId := new(ID)
	fmt.Printf("pId value: %v, pId type: %T\n", pId, pId)

	type Student struct {
		name string
		age  uint8
		ssn  ID
	}
	pStudent := new(Student)
	fmt.Printf("pStudent value: %#v, pStudent type: %T\n", pStudent, pStudent)

	// We can't create pointers from constants or literals
	// const c = 4
	// const s = "Hello, world!"
	// fmt.Printf("address of const c: %v\n", &c)          // illegal
	// fmt.Printf("address of const s: %v\n", &s)          // illegal
	// fmt.Printf(`&true: %v\n`, &true)                    // illegal
	// fmt.Printf(`&1737: %v\n`, &1737)                    // illegal
	// fmt.Printf(`&3.14: %v\n`, &3.14)                    // illegal
	// fmt.Printf(`&complex(11,4): %v\n`, &complex(11, 4)) // illegal
	// fmt.Printf(`&"hello": %v\n`, &"hello")              // illegal
	// type ID string
	// fmt.Printf(`&ID("000-11-2222")`, &ID("000-11-2222"))

	// However, we CAN take pointers from literals of structs
	pStudent2 := &Student{}
	fmt.Printf("pStudent2 value: %#v, address: %p, pStudent2 type: %T\n", pStudent2, pStudent2, pStudent2)

	pStudent2 = &Student{name: "John Doe", age: 45, ssn: "000-00-0000"}
	fmt.Printf("pStudent2 value: %#v, address: %p, pStudent2 type: %T\n", pStudent2, pStudent2, pStudent2)

	// pointer to map literal
	pMap2 := &map[string]string{"name": "John Doe", "age": "45", "ssn": "000-00-0000"}
	fmt.Printf("pMap2 value: %v, address: %p, pMap2 type: %T\n", pMap2, pMap2, pMap2)
}

// Impossible function (for demonstrative purposes only) --> can't implement
func getAddressOfVar(i int) *int {
	return nil
}
