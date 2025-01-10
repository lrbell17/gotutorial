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

	wcm = make(map[string]int)

	var line string
	var fileReaderErr error

	for fileReaderErr == nil {
		line, fileReaderErr = fileReader.ReadLine()
		wcm.processLine(line)
	}
	return
}

func (wcm wordCountMap) processLine(line string) {

	for _, word := range strings.Split(line, " ") {

		if word != "" {
			word = strings.TrimRightFunc(word, trimSpecialChars())
			wcm[word]++
		}
	}
}

func trimSpecialChars() func(rune) bool {
	specialChars := []rune{':', ',', '.', '-'}
	return func(r rune) bool {
		for _, sc := range specialChars {
			if r == sc {
				return true
			}
		}
		return false
	}
}
