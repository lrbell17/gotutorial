package main

import (
	"fmt"
	"os"

	"github.com/lrbell17/gotutorial/sec-05-structs/stub/exercise01/cityutils"
	log "github.com/sirupsen/logrus"
)

const (
	programIdx         = iota
	fileNameIdx        = iota
	numCitiesToCompare = 5
)

func main() {

	fileName := getCmdLineArgs()

	db, err := cityutils.ProcessCsv(fileName)
	if err != nil {
		log.Error("Unable to process CSV file:\n ", err)
		return
	}

	fmt.Printf("Countries with %v or more of the largest cities: %v\n", numCitiesToCompare, db.MostPopulatedCities(numCitiesToCompare))
}

func getCmdLineArgs() string {

	if len(os.Args) <= fileNameIdx {
		usage()
		os.Exit(1)
	}
	return os.Args[fileNameIdx]
}

func usage() {
	log.Errorf("Usage: %v <csv_file>", os.Args[programIdx])
}
