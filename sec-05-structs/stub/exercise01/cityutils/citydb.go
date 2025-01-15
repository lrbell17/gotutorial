package cityutils

import (
	"strconv"
	"strings"

	"github.com/lrbell17/gotutorial/shared/input"
	log "github.com/sirupsen/logrus"
)

type (
	Db map[string][]City
)

const (
	cityNameIdx   = iota
	countryIdx    = iota
	populationIdx = iota
)

func ProcessCsv(fileName string) (Db, error) {

	db := make(Db)

	// Create file reader
	fileReader, fileErr := input.NewFileReader(fileName)
	if fileErr != nil {
		return nil, fileErr
	}

	_, readLineErr := fileReader.ReadLine() // do nothing with line 1 (header)

	for readLineErr == nil {

		// Get line and split by ","
		var line string
		line, readLineErr = fileReader.ReadLine()
		lineArr := strings.Split(line, ",")

		// skip lines with insufficient lines
		if len(lineArr) < 3 {
			log.Warnf("Ignoring line with insufficient fields: \"%v\"", line)
			continue
		}

		// City and country
		cityName := lineArr[cityNameIdx]
		country := lineArr[countryIdx]

		// Population formatted to int
		population64, parseLineErr := strconv.ParseInt(lineArr[populationIdx], 0, 64)
		if parseLineErr != nil {
			return nil, parseLineErr
		}
		population := int(population64)

		// Add record to map
		db[country] = append(db[country], NewCity(cityName, country, population))

	}

	return db, nil

}

func (db Db) MostPopulatedCities(numCities int) (countries []string) {
	for _, cities := range db {
		if len(cities) >= numCities {
			countries = append(countries, cities[0].Country)
		}
	}
	return
}
