package main

import "fmt"

type (
	ID      string
	Student struct {
		name string
		age  uint8
		ssn  ID
	}
)

func main() {

	count := 1737
	pCount := &count
	fmt.Printf("count value: %v, type: %T\n", count, count)
	fmt.Printf("pCount value: %v, type: %T\n", pCount, pCount)

	// The "dereference operator" is "*"
	// used to obtain the value stored at the memory location indicated by a pointer
	fmt.Printf("*pCount value: %v, type: %T\n", *pCount, *pCount)

	// Can derefernce ANY pointer (except for nil pointer)
	p := &Student{"John Doe", 45, "000-00-0000"}
	pp := &p
	fmt.Printf("pp value: %v, type: %T\n", pp, pp)
	fmt.Printf("**pp value: %v, type: %T\n", **pp, **pp)

	// Pitfall - deferencing a nil pointer
	var ps *Student
	fmt.Printf("ps value: %v, type: %T\n", ps, ps)
	// fmt.Printf("*ps value: %v, type: %T\n", *ps, *ps) --> causes panic
	// fmt.Printf("Student name (ps.name): %v\n", ps.name) --> also  causes panic bc we need to dereference p to get the name

	// Good practice to always check for nil defore dereferencing a pointer
	if ps != nil {
		fmt.Printf("*ps value: %v, type: %T\n", *ps, *ps)
		fmt.Printf("Student name (ps.name): %v\n", ps.name)
	}

	// implicit and explicit pointer deference
	// ----
	p = &Student{name: "John Doe", age: 45, ssn: "000-11-2222"}
	fmt.Printf("[DEBUG-1] p's value: %v, type: %T\n", p, p)
	fmt.Printf("[DEBUG-2] *p's value: %v, type: %T\n", *p, *p)

	student := *p
	fmt.Printf("[DEBUG-3] student's value: %v, type: %T\n", student, student)
	fmt.Printf("[DEBUG-4] Student Name (student.name): %v\n", student.name)

	// fmt.Printf("Student Name ((*p).name): %v\n", *p.name) // incorrect, 'Student.name' IS NOT a pointer ("." has higher priority than "*")
	fmt.Printf("[DEBUG-5] Student Name ((*p).name): %v\n", (*p).name) // explicit pointer deference
	fmt.Printf("[DEBUG-6] Student SSN (p.ssn): %v\n", p.ssn)          //  implicit pointer deference (the "." automatically deferences p)

}
