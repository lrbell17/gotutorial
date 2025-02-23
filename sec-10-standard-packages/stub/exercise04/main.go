// Section 10 - Exercise 04 : Simple int/float to string conversion
package main

import (
	"io"
	"os"
)

var (
	digits = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
)

func main() {
	io.WriteString(os.Stdout, intToStr(2010)+"\n")
	io.WriteString(os.Stdout, intToStr(1971/07/07)+"\n") // 40
	io.WriteString(os.Stdout, intToStr(29093423)+"\n")
	io.WriteString(os.Stdout, f64ToStr(9.15)+"\n")
	io.WriteString(os.Stdout, f64ToStr(11.04)+"\n")
}

// intToStr converts an int to a string without any help from any packages
func intToStr(i int) (s string) {
	if i == 0 {
		return "0"
	}

	// convert each digit to string
	for i != 0 {
		s += string(i%10 + '0')
		i = i / 10
	}

	// reverse order
	r := []rune(s)
	for x, y := 0, len(s)-1; x < y; x, y = x+1, y-1 {
		r[x], r[y] = r[y], r[x]
	}

	return string(r)
}

// f64ToStr converts a float64 to a string without any help from any packages
func f64ToStr(f float64) (s string) {

	whole := int(f)

	s = intToStr(whole) + "."

	next := f - float64(whole)

	for i := 0; i < 3; i++ {
		dec := int(next * 10)
		s += intToStr(dec)
		next = next*10 - float64(dec)
	}

	return
}
