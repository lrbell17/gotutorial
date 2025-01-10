package main

import (
	"fmt"

	"github.com/lrbell17/gotutorial/sec-02-basic-concepts/stub/lab01/dbutils"
)

func main() {
	fmt.Println("Lowest income: ", dbutils.GetLowestIncome(dbutils.GetRecord()))
	fmt.Println("Highest income: ", dbutils.GetHighestIncome(dbutils.GetRecord()))
	fmt.Println("Average income: ", dbutils.GetAvgIncome(dbutils.GetRecord()))

	fmt.Println("Lowest income (male): ", dbutils.GetLowestIncome(dbutils.GetRecordByGender("Male")))
	fmt.Println("Highest income (male): ", dbutils.GetHighestIncome(dbutils.GetRecordByGender("Male")))
	fmt.Println("Average income (male): ", dbutils.GetAvgIncome(dbutils.GetRecordByGender("Male")))

	fmt.Println("Lowest income (female): ", dbutils.GetLowestIncome(dbutils.GetRecordByGender("Female")))
	fmt.Println("Highest income (female): ", dbutils.GetHighestIncome(dbutils.GetRecordByGender("Female")))
	fmt.Println("Average income (female): ", dbutils.GetAvgIncome(dbutils.GetRecordByGender("Female")))
}
