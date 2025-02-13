package main

import "fmt"

type (
	ID     uint64
	SSN    string
	Person struct {
		name string
		age  uint8
		ssn  SSN
	}
	Printer interface {
		Print() string
	}
)

func main() {
	p := getValue()
	Println(p)
}

// can return any type that implements Printer
func getValue() Printer {

	ch := make(chan int)

	go func() {
		// either 0 or 1 with equal probability
		select {
		case ch <- 0:
		case ch <- 1:
		}

	}()
	choice := <-ch

	if choice == 0 {
		return SSN("000-00-0000")
	}

	return ID(123456)
}

func (id ID) Print() string {
	return fmt.Sprintf("[ID] %v", uint64(id))
}

func (ssn SSN) Print() string {
	return fmt.Sprintf("[SSN] %v", string(ssn))
}

func Println(e ...Printer) {
	fmt.Println("[main.Println]")
	for i := range e {
		if e[i] == nil {
			fmt.Printf("Parameter %v is nil\n", i)
			continue
		}
		fmt.Printf("Parameter %v's .Print() value: %v\n", i, e[i].Print())
	}
}
