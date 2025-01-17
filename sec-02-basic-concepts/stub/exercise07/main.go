/*
	IMPORTANT: You *must* use 'continue' in at least one of

your solutions.
---------------------------------------------------------------
TODO 1 - calculate the sum of ODD nubmers between 1 to 10001
HINT: Use the modulo operator:

	https://www.khanacademy.org/computing/computer-science/cryptography/modarithmetic/a/what-is-modular-arithmetic

ODD numbers  -> 1, 3, 5, 7, 9, 11, .... 9999, 10001
Result should be: 25010001

TODO 2 - calculate the 'integer' average of all nubmers from 1 to 10001,
excluding the nubmers:

	10, 19, 21, 39, 309, 431, 643, 942, 1209, 7981, 8888, 9910

Result should be: 5003
*/
package main

import "fmt"

func main() {
	var oddSum int
	var totalSum int
	var totalCount int
	for i := 1; i <= 10001; i++ {
		if i%2 != 0 {
			oddSum += i
		}

		switch i {
		case 10, 19, 21, 39, 309, 431, 643, 942, 1209, 7981, 8888, 9910:
			continue
		}
		totalSum += i
		totalCount++
	}
	fmt.Printf("Sum of odd numbers: %v\n", oddSum)
	fmt.Printf("Avg: %v\n", totalSum/totalCount)
}
