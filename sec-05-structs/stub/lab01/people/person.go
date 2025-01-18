package people

import "fmt"

type (
	People []Person
	Person struct {
		firstName string
		lastName  string
		ssn       string
		gender    string
		age       int8
		salary    Curreny
	}
	Curreny float32
)

func (c Curreny) String() (s string) {
	s = fmt.Sprintf("$%.2f", float32(c))
	return
}
