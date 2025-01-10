package dbutils

import (
	"github.com/lrbell17/gotutorial/shared/sec02/db"
	log "github.com/sirupsen/logrus"
)

func init() {
	err := db.Load()
	if err != nil {
		log.Fatal("Could not load DB")
	}
}

func GetHighestIncome(recordGetter RecordGetter) (employee string) {

	db.ResetCursor()
	var highestIncome db.Currency

	for {
		name, genderMatch, income, err := recordGetter()

		if genderMatch && income > highestIncome {
			highestIncome = income
			employee = name
		}

		if err != nil {
			break
		}

	}
	return
}

func GetLowestIncome(recordGetter RecordGetter) (employee string) {
	db.ResetCursor()
	var lowestIncome db.Currency = -1

	for {
		name, genderMatch, income, err := recordGetter()

		if genderMatch && (income < lowestIncome || lowestIncome < 0) {
			lowestIncome = income
			employee = name
		}

		if err != nil {
			break
		}

	}
	return
}

func GetAvgIncome(recordGetter RecordGetter) (avg db.Currency) {
	db.ResetCursor()
	totalEmployees := 0
	totalIncome := db.Currency(0.0)

	for {
		_, genderMatch, income, err := recordGetter()

		if genderMatch {
			totalEmployees++
			totalIncome += income
		}

		if err != nil {
			break
		}

	}
	return totalIncome / db.Currency(totalEmployees)
}
