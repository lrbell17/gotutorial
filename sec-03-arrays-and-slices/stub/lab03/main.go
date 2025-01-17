package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lrbell17/gotutorial/sec-03-arrays-and-slices/stub/lab03/grades"
)

const (
	numStudentsIdx = iota + 1
	numExamsIdx    = iota + 1
	maxScoreIdx    = iota + 1
	classNameIdx   = iota + 1
)

func main() {
	_, _, _, className := getCmdLineArgs()

	classReport := grades.BuildClassReport(getCmdLineArgs())

	fmt.Printf("\nClass report card for: %v\n", className)
	printReport(classReport)

}

func printReport(classReport grades.ClassReport) {
	testReport := classReport.GetTestReport()

	printHeader(len(testReport))
	printScores(classReport)
	printAvg(testReport)

}

func printHeader(numTests int) {
	fmt.Println(strings.Repeat("-", 15*(2+numTests)))

	fmt.Printf("%-15s", "Student ID")
	for i := 1; i <= numTests; i++ {
		fmt.Printf("%-15s", fmt.Sprintf("Test %d", i))
	}
	fmt.Println("Letter Grade")

	fmt.Println(strings.Repeat("-", 15*(2+numTests)))

}

func printScores(classReport grades.ClassReport) {

	for i, studentReport := range classReport {
		fmt.Printf("%-15d", i+1)
		for _, score := range studentReport {
			fmt.Printf("%-15d", score)
		}
		fmt.Printf("%-15s\n", studentReport.GetStudentGrade())
	}
}

func printAvg(testReport grades.TestReport) {

	fmt.Println(strings.Repeat("-", 15*(2+len(testReport))))
	fmt.Printf("%-15s", "Average:")

	for _, v := range testReport {
		fmt.Printf("%-15d", v)
	}
	fmt.Println()
}

func getCmdLineArgs() (int, int, int, string) {

	numStudents, parseErr1 := strconv.ParseInt(os.Args[numStudentsIdx], 0, 32)
	numExams, parseErr2 := strconv.ParseInt(os.Args[numExamsIdx], 0, 32)
	maxScore, parseErr3 := strconv.ParseInt(os.Args[maxScoreIdx], 0, 32)
	className := os.Args[classNameIdx]

	if parseErr1 != nil || parseErr2 != nil || parseErr3 != nil {
		fmt.Println("Error parsing arguments")
		os.Exit(1)
	}

	if numStudents < 1 {
		fmt.Println("The number of students must be > 0")
		os.Exit(1)
	}

	if numExams < 1 {
		fmt.Println("The number of exams must be > 0")
		os.Exit(1)
	}

	if maxScore < 100 || maxScore > 120 {
		fmt.Println("The max score must be between 100 and 120")
		os.Exit(1)
	}

	return int(numStudents), int(numExams), int(maxScore), className
}
