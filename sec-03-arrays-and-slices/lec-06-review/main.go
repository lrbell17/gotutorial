package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/lrbell17/gotutorial/shared/input"
)

func main() {

	// ---------------Program arguments ----------------------------
	fmt.Printf("os.Args type: %T\n", os.Args)

	/*
		Note: get different results with 2 scenarios:

		go run main.go
			os.Args[0] = /var/folders/dj/8g0_lnk56t399rpnhv23p2500000gp/T/go-build4106765664/b001/exe/main

		go build main.go; ./main
			os.Args[0] = ./main
	*/
	for i, v := range os.Args {
		fmt.Printf("os.Args[%v] = %v\n", i, v)
	}

	/*
		lukebell@managers-MacBook-Pro lec-06-review % ./awesome hello world
		os.Args type: []string
		os.Args[0] = ./awesome
		os.Args[1] = hello
		os.Args[2] = world

		// double quotes:
		lukebell@managers-MacBook-Pro lec-06-review % ./awesome "hello world - ${GOPATH}"
		os.Args type: []string
		os.Args[0] = ./awesome
		os.Args[1] = hello world - /Users/lukebell/go

		// single quotes - literal
		lukebell@managers-MacBook-Pro lec-06-review % ./awesome 'hello world - ${GOPATH}'
		os.Args type: []string
		os.Args[0] = ./awesome
		os.Args[1] = hello world - ${GOPATH}
	*/

	//  ./main 3 3.14 true
	if len(os.Args) > 1 {
		intVal, _ := strconv.ParseInt(os.Args[1], 0, 64)    // base (0, 2 - 36), bitSize(0-64)
		floatVal, err := strconv.ParseFloat(os.Args[2], 32) // pass 32 or 64 bits
		boolVal, _ := strconv.ParseBool(os.Args[3])         // works with 0 and 1

		fmt.Println(err)

		fmt.Printf("intVal: %v, type: %T\n", intVal, intVal)
		fmt.Printf("floatVal: %v, type: %T\n", floatVal, floatVal) // returns only an approximate value!!!
		fmt.Printf("boolVal: %v, type: %T\n", boolVal, boolVal)
	}

	// ------------ Multi-dimensional Arrays/Slices -------------------------

	table1, _ := initTable(4, 7, 3, 5, 4)
	printTable(table1)

	_, err := initTable(4, 7, 3, 5)
	fmt.Println(err)

	// ----------------- Re-slicing a slice --------------------------------

	nums := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s1 := nums[:5]
	fmt.Printf("s1: %v, len: %v\n", s1, len(s1))

	s2 := s1[3:8]
	fmt.Printf("s2: %v, len: %v\n", s2, len(s2)) // returns values from the original array not in s1!!!

	s3 := s1[:]
	fmt.Printf("s3: %v, len: %v\n", s3, len(s3)) // just returns same as s1

	// ----------------------- Strings -------------------------------------

	// a string is an array of runes/bytes

	s := "Hello, 世界"                             // 9 chars
	fmt.Printf("s: %v, len(s): %v\n", s, len(s)) // returns 13 since chinese chars are 3 bytes

	// returns type int8 for each byte
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%v]: %v, type: %T\n", i, s[i], s[i])
	}

	// returns type rune (int32)
	for i, v := range s {
		fmt.Printf("s[%v]: %v, rune val: \"%c\", type: %T\n", i, v, v, v)
	}

	// string as []byte
	sab := []byte(s)
	fmt.Printf("sab: %v, len(sab): %v\n", sab, len(sab))

	// string as []rune
	sar := []rune(s)
	fmt.Printf("sar: %v, len(sar): %v\n", sar, len(sar))

	// cast to string
	str2 := string(sab)
	str3 := string(sar)
	fmt.Printf("str2: %v, len(str2): %v\n", str2, len(str2))
	fmt.Printf("str3: %v, len(str3): %v\n", str3, len(str3))

}

type Row []int
type Table []Row

func initTable(rows int, cols ...int) (Table, error) {

	if rows != len(cols) {
		return nil, errors.New("invalid column length")
	}

	var table Table = make(Table, rows)

	for i := range table {
		table[i] = make(Row, cols[i])
		for j := range table[i] {
			table[i][j] = input.GetRandInt(-20, 120)
		}
	}
	return table, nil
}

func printTable(table Table) {
	for i := range table {
		for j := range table[i] {
			fmt.Printf("%5v", table[i][j]) // print w/ padding
		}
		fmt.Println()
	}
}
