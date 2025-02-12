package main

import "fmt"

/*
A set of methods can be selected for a value or interface variable

  - if v is a value of type T or *T, v's method set is all the methods callable by a selector on v, such as v.M()

  - if i is an interface variable of the interface type I, i's method set is all the methods defined for interface I, such as i.M()

Golang specs:

  - A type *may* have a method set assocciated with it

  - A method set of an interface type is its interface

  - A method set of any type T consists of all methods declared with receiver type T

  - A method set of the corresponding pointer type *T is the set of methods declared with recceiver type T or *T
    --> i.e. it also contains method set of T

  - Any other type has an empty method set
*/

/*
--- Method set of T and *T ---

type Person struct{}
methods:

	func (p Person) Name() string
	func (p *Person) SetAge(a uint8)

--> method set of T is Name()
--> method set of *T is Name(), SetAge()
*/

type (
	Person struct {
		fname string
		lname string
		age   uint8
	}
)

func main() {
	// Method sets for T - Name(), Age()
	mark := Person{"Mark", "Smith", 35}
	fmt.Println(mark.Name())
	fmt.Printf("Mark's age: %v\n", mark.Age())

	// Method sets for *T - Name(), Age(), SetAge()

	// using variable of type T to access *T method set
	mark.SetAge(mark.Age() + 1) // same as m:= &mark; m.SetAge(...)
	fmt.Printf("Mark's age: %v\n", mark.Age())

	// method set of *T also includes method set of T
	p := &Person{"Jane", "Doe", 39}
	p.SetAge(p.Age() + 1)
	fmt.Println(p.Name()) // same as *p.Name()

	fmt.Println()

	// However, the method set of T does not contain the method set of *T
	family := map[string]Person{
		"mom": {"Jane", "Smith", 42},
		"dad": {"John", "Smith", 45},
	}

	// the following doesn't work since only a copy is in the map, can't get the pointer
	for r, p := range family {
		fmt.Printf("Role: %v, Name: %v, Age: %v\n", r, p.Name(), p.Age())
	}
	for _, p := range family {
		fmt.Println(&p)
		p.SetAge(p.Age() + 1)
	}
	for r, p := range family {
		fmt.Printf("Role: %v, Name: %v, Age: %v\n", r, p.Name(), p.Age())
	}

}

func (p Person) Name() string {
	return fmt.Sprintf("%v, %v", p.lname, p.fname)
}

func (p Person) Age() uint8 {
	return p.age
}

func (p *Person) SetAge(a uint8) {
	if a < 150 && a > p.age {
		fmt.Printf("Changing age of %v from %v to %v\n", p.Name(), p.age, a)
		p.age = a
	}
}
