package people

import (
	"errors"
	"strconv"
	"strings"

	"github.com/lrbell17/gotutorial/shared/input"
	log "github.com/sirupsen/logrus"
)

const (
	firstNameIdx = iota
	lastNameIdx  = iota
	ssnIdx       = iota
	genderIdx    = iota
	ageIdx       = iota
	salaryIdx    = iota
)

func ReadPersonData(fileName string) (People, error) {

	fileReader, invalidFileErr := input.NewFileReader(fileName)
	if invalidFileErr != nil {
		log.Error(invalidFileErr)
		return nil, invalidFileErr
	}

	var people People

	var eofErr error
	var rawLine string
	for eofErr == nil {

		rawLine, eofErr = fileReader.ReadLine()
		if strings.TrimSpace(rawLine) == "" {
			log.Warn("Empty line, skipping")
			continue
		}

		line := strings.Split(rawLine, ",")
		gender, age, salary, parseErr := parsePersonDetails(line)
		if parseErr != nil {
			log.Warnf("%v\nThe offending line was: %v\n", parseErr, line)
			continue
		}

		person := Person{
			firstName: line[firstNameIdx],
			lastName:  line[lastNameIdx],
			ssn:       line[ssnIdx],
			gender:    gender,
			age:       age,
			salary:    salary,
		}

		people = append(people, person)

	}

	return people, nil
}

func parsePersonDetails(line []string) (string, int8, Curreny, error) {
	gender, parseGenderErr := parseGender(line[genderIdx])
	if parseGenderErr != nil {
		return "", 0, 0, parseGenderErr
	}

	age, parseAgeErr := parseAge(line[ageIdx])
	if parseAgeErr != nil {
		return "", 0, 0, parseAgeErr
	}

	salary, parseSalaryErr := parseSalary(line[salaryIdx])
	if parseSalaryErr != nil {
		return "", 0, 0, parseSalaryErr
	}

	return gender, age, salary, nil
}

func parseAge(ageStr string) (age int8, parseErr error) {

	parsedAge, parseErr := strconv.ParseInt(ageStr, 0, 8)
	if parseErr != nil {
		return 0, parseErr
	}
	age = int8(parsedAge)

	if age < 0 {
		return 0, errors.New("the age must be greater than 0")
	}
	return
}

func parseSalary(salaryStr string) (salary Curreny, err error) {

	parsedSalary, parseErr := strconv.ParseFloat(salaryStr, 32)
	if parseErr != nil {
		return 0, parseErr
	}
	salary = Curreny(parsedSalary)

	if salary < 0 {
		return 0, errors.New("the salary must be greater than 0")
	}
	return
}

func parseGender(genderStr string) (string, error) {

	if genderStr != "M" && genderStr != "F" {
		return "", errors.New("invalid gender, must be M or F")
	}
	return genderStr, nil
}
