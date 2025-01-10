package dbutils

import "github.com/lrbell17/gotutorial/shared/sec02/db"

type RecordGetter = func() (name string, genderMatch bool, income db.Currency, err error)

func GetRecord() RecordGetter {

	return func() (name string, genderMatch bool, income db.Currency, err error) {
		_, firstName, lastName, _, income, err := db.GetNextRecord()
		name = firstName + " " + lastName
		return name, true, income, err
	}

}

func GetRecordByGender(targetGender string) RecordGetter {

	return func() (name string, genderMatch bool, income db.Currency, err error) {
		_, firstName, lastName, gender, income, err := db.GetNextRecord()
		name = firstName + " " + lastName
		if targetGender == gender {
			genderMatch = true
		}
		return name, genderMatch, income, err
	}
}
