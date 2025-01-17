package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/lrbell17/gotutorial/shared/input"
	log "github.com/sirupsen/logrus"
)

const (
	timeIdx     = iota + 1
	operatorIdx = iota + 1
	durationIdx = iota + 1

	currentTimeVal = "now"
	operatorVal    = "+"
)

func main() {
	_, _, duration, err := getCmdLineArgs()
	if err != nil {
		log.Error(err)
		return
	}

	reminder := input.PromptString("Enter your reminder: ")

	time.Sleep(duration)
	fmt.Println(reminder)
}

func getCmdLineArgs() (currentTime string, operator string, duration time.Duration, err error) {

	if len(os.Args) != 4 {
		return "", "", 0, errors.New("invalid number of arguments")
	}

	currentTime = os.Args[timeIdx]
	operator = os.Args[operatorIdx]
	durationAsStr := os.Args[durationIdx]

	duration, err = time.ParseDuration(durationAsStr)

	if currentTime != currentTimeVal {
		return "", "", 0, fmt.Errorf("the current time must be \"%v\"", currentTimeVal)
	}

	if operator != operatorVal {
		return "", "", 0, fmt.Errorf("the operator must be \"%v\"", operatorVal)
	}

	return currentTime, operator, duration, err

}
