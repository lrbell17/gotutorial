package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const maxLines = 5

func main() {
	var buf []string
	line := 1
	fmt.Printf("Enter at most %v lines of text below (empty line to quit)\n", maxLines)
	fmt.Printf("#%v>", line)

	scanner := bufio.NewScanner(os.Stdin)

	// reading lines from input
	for line <= maxLines && scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			break
		}

		buf = append(buf, s)
		line++
		fmt.Printf("#%v>", line)
	}

	// print lines
	fmt.Println("\nOutput:")
	for i := range buf {
		s := strings.Title(buf[i]) // strings.Title capitalizes first letter of each word
		fmt.Printf("#%v: %v\n", i+1, s)
	}

}
