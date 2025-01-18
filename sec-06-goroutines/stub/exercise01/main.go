package main

import (
	"fmt"
	"time"
)

const (
	capacity = 100000
	workers  = 6
)

var nums []int

func main() {

	for i := 0; i < workers; i++ {
		go func() {
			slice := subSlice(nums[:], workers, i)
			primeSearch(i+1, slice)
		}()
	}

	time.Sleep(5 * time.Second)
}

func init() {
	nums = make([]int, capacity)
	for i := range nums {
		nums[i] = i + 1
	}
}

func primeSearch(workerId int, slice []int) {

	for i := 1; i < len(slice); i++ {

		v := slice[i]
		prime := true

		for j := 2; j < v/2+1; j++ {
			if v%j == 0 {
				prime = false
				break
			}
		}
		if prime {
			fmt.Println(v)
		}
	}
}
