package main

import "fmt"

/*
Interface variable has 2 hidden fields:
- <value> holds a copy of the assigned value
- <type> hold info about the value's type

Empty interface: var foo interface{}
- value: nil
- type: nil

Assign value to interface: foo = 3.14
- value: 3.14
- type: float64

Assign another value to interface: foo = &Person{}
- value: 0x11face04
- type: *Person
*/

type (
	ID     uint64
	SSN    string
	Person struct {
		name string
		age  uint8
		ssn  SSN
	}
	Empty   interface{}
	Printer interface {
		Print() string
	}
)

func main() {

	var e Empty
	printInfo(e)
	printInfo(11.04)
	printInfo(ID(123456))
	printInfo(SSN("000-00-0000"))
	printInfo(&Person{"John Doe", 35, SSN("000-00-0001")})

	// variadic:
	Println(e, 11.04, ID(213456), SSN("000-00-0000"), &Person{"John Doe", 35, SSN("000-00-0001")})

	// printer interface - all types need to implement the Print function
	var p Printer
	PrinterPrint(p, ID(213456), SSN("000-00-0000"), &Person{"John Doe", 35, SSN("000-00-0001")})

}

// since we are using an empty interface, we can use any type here!
func printInfo(e Empty) {
	fmt.Printf("e's value: %v, Type: %T\n", e, e)
}

// any is an alias for empty interface
func Println(a ...any) {
	fmt.Println("[main.Println]")
	for _, v := range a {
		fmt.Printf("e's value: %v, Type: %T\n", v, v)
	}
}

func PrinterPrint(a ...Printer) {
	fmt.Println("[main.PrinterPrint]")
	for i, v := range a {
		if v == nil {
			fmt.Printf("Parameter[%v] is nil\n", i)
			continue
		}
		fmt.Printf("Value: %v, Type: %T\n", v, v)
	}
}

// This doesn't work - can't define methods on a non-locacl type
// func (f float64) Print() string {
// 	return fmt.Sprintf("[float64] %v", f)
// }

func (ssn SSN) Print() string {
	return fmt.Sprintf("[SSN] %v", string(ssn))
}

func (id ID) Print() string {
	return fmt.Sprintf("[ID] %v", uint8(id))
}

func (p *Person) Print() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("[Person] Name: %v, Age: %v, SSN: %v", p.name, p.age, p.ssn)
}
