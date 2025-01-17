package main

import "fmt"

type (
	Address struct {
		line1 string
		line2 string
	}
	Age    uint8
	Person struct {
		firstName string
		lastName  string
		ssn       string
		age       Age
		address   Address
	}
	// Anonymous fields
	City struct {
		string
		// string -> this is illegal since string is already used
		State
	}
	State struct {
		name string
	}
	Country struct {
		name string
	}
	City2 struct {
		string
		State
		Country
	}
)

func main() {

	// initializing
	var a0 Address
	a0.line1 = "10 6th Ave"
	a0.line2 = "Apt B"
	fmt.Printf("%#v\n", a0) // use # to include field names

	var a1 = Address{line1: "10 6th Ave", line2: "Apt B"} // can initialize all or some of the fields
	fmt.Printf("%#v\n", a1)

	var a2 = Address{"10 6th Ave", "Apt B"} // don't need to include field names if ordering is preseved
	fmt.Printf("%#v\n", a2)

	var a3 = Address{
		line1: "10 6th Ave",
		line2: "Apt B",
	}
	fmt.Printf("%#v\n", a3)

	// Nested structs
	p0 := Person{
		firstName: "John",
		lastName:  "Doe",
		ssn:       "000-00-0000",
		age:       18,
		address: Address{
			line1: "123 Atlantic Ave",
			line2: "Suite 321",
		},
	}
	fmt.Printf("%#v\n", p0)

	sam := Person{
		"Sam",
		"Smith",
		"000-00-0001",
		29,
		Address{"567 Pacific St", "Apt 890"},
	}
	fmt.Printf("%#v\n", sam)

	// Anonymous fields
	c := City{"Brooklyn", State{"New York"}}
	fmt.Printf("City: %v, State: %v\n", c.string, c.name) // c.name is shorthand for c.State.name
	fmt.Printf("City: %v, State: %v\n", c.string, c.State.name)

	// ambiguos fields
	c2 := City2{State: State{"New York"}, string: "Brooklyn", Country: Country{"USA"}}
	// fmt.Printf("City: %v, State: %v, Country:%v\n", c2.string, c2.name, c2.name) --> not allowed due to ambiguous selector
	fmt.Printf("City: %v, State: %v\n", c.string, c2.State.name)
}
