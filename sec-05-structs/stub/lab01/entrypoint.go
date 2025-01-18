package main

import (
	"os"

	"github.com/lrbell17/gotutorial/sec-05-structs/stub/lab01/people"
	log "github.com/sirupsen/logrus"
)

const (
	programIdx  = iota
	fileNameIdx = iota
)

func main() {

	if len(os.Args) != fileNameIdx+1 {
		usage()
		return
	}
	fileName := os.Args[fileNameIdx]

	people, err := people.ReadPersonData(fileName)
	if err != nil {
		usage()
		return
	}

	people.GetStatsForGender("M")
	people.GetStatsForGender("F")

}

func usage() {
	log.Errorf("Usage: %v <filename>\n", os.Args[programIdx])
}
