package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/lrbell17/gotutorial/shared/input"

	log "github.com/sirupsen/logrus"
)

const (
	fileNameIdx = 1
)

type wordCountMap map[string]int

func main() {

	fileName := os.Args[fileNameIdx]

	if fileReader, err := input.NewFileReader(fileName); err != nil {
		log.Error(err)
		return
	} else {
		wcm := countWordsInFile(*fileReader)
		for k, v := range wcm {
			fmt.Printf("%v:  %v\n", k, v)
		}
	}

}

func countWordsInFile(fileReader input.FileReader) (wcm wordCountMap) {

	specialChars := []rune{':', ',', '.', '-'}
	wordsToIgnore := []string{"on", "a", "the", "are", "in", "of", "for"}

	wcm = make(map[string]int)

	var line string
	var fileReaderErr error

	for fileReaderErr == nil {
		line, fileReaderErr = fileReader.ReadLine()
		wcm.processLine(line, specialChars, wordsToIgnore)
	}
	return
}

func (wcm wordCountMap) processLine(line string, specialChars []rune, wordsToIgnore []string) {

	for _, word := range strings.Split(line, " ") {

		if !contains(wordsToIgnore, word) {
			word = strings.TrimRightFunc(word, func(r rune) bool {
				return contains(specialChars, r)
			})
			wcm[word]++
		}
	}
}

func contains[T comparable](slice []T, str T) bool {

	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
