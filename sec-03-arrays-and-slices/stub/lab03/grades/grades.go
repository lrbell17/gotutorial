package grades

import (
	"math/rand"
)

const (
	minScore = 30
)

type (
	StudentReport []int
	ClassReport   []StudentReport
	TestReport    []int
)

func BuildClassReport(numStudents int, numExams int, maxScore int, className string) ClassReport {

	classReport := make(ClassReport, numStudents)

	for i := range classReport {
		classReport[i] = buildStudentReport(numExams, maxScore)
	}

	return classReport

}

func (classReport ClassReport) GetTestReport() (testScores TestReport) {

	numStudents := len(classReport)
	numTests := len(classReport[0])
	testScores = make([]int, numTests)

	for testNum := 0; testNum < numTests; testNum++ {
		var totalScore int
		for _, studentReport := range classReport {
			totalScore += studentReport[testNum]
		}
		testScores[testNum] = totalScore / numStudents
	}

	return
}

func buildStudentReport(numExams int, maxScore int) StudentReport {
	studentReport := make(StudentReport, numExams)

	for i := range studentReport {
		studentReport[i] = minScore + rand.Intn(maxScore-minScore+1)
	}

	return studentReport
}

func (studentReport StudentReport) GetStudentGrade() (grade string) {

	var totalScore int
	for _, v := range studentReport {
		totalScore += v
	}

	avgScore := totalScore / len(studentReport)

	return mapScoreToGrade(avgScore)

}

func mapScoreToGrade(score int) string {

	scoreMap := []int{98, 95, 90, 80, 70, 60, 55, 45}
	gradeMap := []string{"A+", "A", "B+", "B", "C+", "C", "D+", "D"}

	for i := range scoreMap {
		if score >= scoreMap[i] {
			return gradeMap[i]
		}
	}
	return "F"
}
