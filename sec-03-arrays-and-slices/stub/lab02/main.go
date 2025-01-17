package main

import (
	"errors"
	"fmt"

	"github.com/lrbell17/gotutorial/shared/input"
)

func main() {

	err := userPrompt()

	if err != nil {
		fmt.Println("Error: ", err)
	}

}

func userPrompt() error {

	numScores := input.PromptInt("Enter the number of test scores: ")
	if numScores < 0 {
		return errors.New("the number of scores must be >= 0")
	}

	scores := make([]int, numScores)

	for i := range scores {
		score := input.PromptInt(fmt.Sprintf("score %v: ", i+1))
		if score < 0 || score > 100 {
			return errors.New("the score must be between 0 and 100")
		} else {
			scores[i] = score
		}
	}

	fmt.Printf("%-10s %-10s %-10s\n", "Test #", "Score", "Grade")
	for i, v := range scores {
		fmt.Printf("%-10d %-10d %-10s\n", i+1, v, calculateGrade(v))
	}

	return nil

}

func calculateGrade(score int) (grade string) {

	switch {
	case score >= 98:
		grade = "A+"
	case score >= 95:
		grade = "A"
	case score >= 90:
		grade = "B+"
	case score >= 80:
		grade = "B"
	case score >= 70:
		grade = "C+"
	case score >= 60:
		grade = "C"
	case score >= 55:
		grade = "D+"
	case score >= 45:
		grade = "D"
	default:
		grade = "F"
	}

	return
}
