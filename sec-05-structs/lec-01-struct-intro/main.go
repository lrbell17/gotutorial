package main

import "fmt"

/*
	A struct is a sequence of named elements (fields), each of which have a name and type
	- Field names may be specified explicitly (IdentifierList) or implcitly (AnonymousField)
	- Non-blank field names must be unique

	Arrays vs. Slices vs. Maps vs. Structs
	- Arrays:
		--> flexible element type (all elements must be the same type)
		--> fixed boundaries
		--> non-neg integer indexing
	- Slices:
		--> same as Array, but resizeable
	- Maps:
		--> like Slices, but with flexible key (all keys muse be same type)
	- Structs:
		--> has fields that can be different types
*/

type (
	PersonHack []string
	FamilyHack []PersonHack
	DbHack     map[LastName]FamilyHack
	LastName   string

	Person struct {
		firstName string
		lastName  string
		ssn       string
		age       int
		// height      float32
		// middleNames []string
		// phones      map[string]string
	}
	Family []Person
	Db     map[LastName]Family
)

func main() {

	// without structs
	db := make(DbHack)
	mary := newPersonHack("Mary", "Smith", "000-00-0001", "21")
	john := newPersonHack("John", "Smith", "000-00-0002", "24")
	pete := newPersonHack("Pete", "Smith", "000-00-0003", "45")
	anne := newPersonHack("Anne", "Smith", "000-00-0004", "63")
	db["Smith"] = append(db["Smith"], mary, john, pete, anne)
	fmt.Printf("Smith family: %v\n", db["Smith"])

	// Use a struct
	var max Person
	max.firstName = "Max"
	max.lastName = "Jones"
	max.ssn = "001-00-0001"
	max.age = 30
	fmt.Printf("Max: %v, type: %T\n", max, max)

	// create DB with struct
	db2 := make(Db)
	mary2 := newPerson("Mary", "Smith", "000-00-0001", 21)
	john2 := newPerson("John", "Smith", "000-00-0002", 24)
	pete2 := newPerson("Pete", "Smith", "000-00-0003", 45)
	anne2 := newPerson("Anne", "Smith", "000-00-0004", 63)

	db2["Smith"] = append(db2["Smith"], mary2, john2, pete2, anne2)
	db2["Jones"] = append(db2["Jones"], max)
	fmt.Printf("db2: %v, Type: %T\n", db2, db2)

}

func newPersonHack(fn, ln, ssn, age string) (p PersonHack) {
	p = append(p, fn, ln, ssn, age)
	return
}

func newPerson(fn, ln, ssn string, age int) (p Person) {
	p.firstName = fn
	p.lastName = ln
	p.ssn = ssn
	p.age = age
	return
}
