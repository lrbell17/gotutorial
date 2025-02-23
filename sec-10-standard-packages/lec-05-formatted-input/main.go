package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// scan into from a string
	input := "41"

	var a int
	fmt.Sscan(input, &a) // give it a pointer to a to store result
	fmt.Printf("The answer the to universe is: %v\n", a+1)

	// scan int, float, and strings from a string
	input = "Jane Doe 35 001-11-2233 5.6"
	var i int
	var f float64
	var s0, s1, s2 string
	fmt.Sscan(input, &s0, &s1, &i, &s2, &f) // scans separated by whitespace
	fmt.Printf("\nGot int: %v\n", i)
	fmt.Printf("Got float64: %v\n", f)
	fmt.Printf("Got string s0: %v\n", s0)
	fmt.Printf("Got string s1: %v\n", s1)
	fmt.Printf("Got string s2: %v\n\n", s2)

	// Scanln will terminate when there is a newline
	input2 := "Jane Doe\n35\n001-11-2233\n5.6"
	var ii int
	var ff float64
	var s00, s11, s22 string
	fmt.Sscanln(input2, &s00, &s11, &ii, &s22, &ff)
	fmt.Printf("\nGot int: %v\n", ii)
	fmt.Printf("Got float64: %v\n", ff)
	fmt.Printf("Got string s00: %v\n", s00)
	fmt.Printf("Got string s11: %v\n", s11)
	fmt.Printf("Got string s22: %v\n\n", s22)

	// reading file
	readFile()

	// reading from standard input
	readStdin()
}

func readFile() {
	input, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	var i int
	var f float64
	var s0, s1 string

	fmt.Fscanln(input, &s0, &i, &s1, &f)
	fmt.Printf("\nGot int: %v\n", i)
	fmt.Printf("Got float64: %v\n", f)
	fmt.Printf("Got string s0: %v\n", s0)
	fmt.Printf("Got string s1: %v\n", s1)
}

func readStdin() {

	fmt.Println("\nEnter input: ")
	input := os.Stdin // we don't need to close

	var i int
	var f float64
	var s0, s1 string

	fmt.Fscanln(input, &s0, &i, &s1, &f)
	fmt.Printf("\nGot int: %v\n", i)
	fmt.Printf("Got float64: %v\n", f)
	fmt.Printf("Got string s0: %v\n", s0)
	fmt.Printf("Got string s1: %v\n", s1)
}
