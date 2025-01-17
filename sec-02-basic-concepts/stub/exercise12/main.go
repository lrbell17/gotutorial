package main

import (
	"fmt"

	"github.com/lrbell17/gotutorial/shared/cal"
	"github.com/lrbell17/gotutorial/shared/input"
)

func main() {
	fmt.Print("Enter your name: ")
	name := input.ReadString()
	date := cal.DateNow()
	time := cal.TimeNow()

	fmt.Printf("Hello %v\n", name)
	fmt.Printf("Today is %v. The current time is %v\n", date, time)
}
