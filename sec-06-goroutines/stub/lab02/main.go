// Section 06 - Lab 01 : Work Count (iterative)
package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/lrbell17/gotutorial/shared/input"
	log "github.com/sirupsen/logrus"
)

var (
	result       map[string]int // word(string) -> count(int)
	specialChars []string

	wg sync.WaitGroup
	mt sync.Mutex
)

func init() {
	specialChars = []string{"\"", ",", ".", "(", ")", "!", "?", "-", "]", "[", "'", ";", "*", "_"}
}

func main() {

	if 1 == len(os.Args) {
		log.Fatal("No files to process")
	}

	inputFiles := os.Args[1:]
	result = make(map[string]int)

	start := time.Now()
	log.Infof("Processing %v files for input", len(inputFiles))
	for _, fn := range inputFiles {
		processFile(fn)
	}

	wg.Wait()

	elapse := time.Since(start)
	printResult()
	fmt.Printf("Total time: %v\n", elapse)
}

// processFile takes a filename 'fn' and reads each line of text. For
// each line, it call countWords() to count the number of words in the line.
func processFile(fn string) {
	log.Infof("Processing file %v", fn)

	wg.Add(1)
	go func() {

		fr, frErr := input.NewFileReader(fn)
		if frErr != nil {
			log.Fatalf("Could not read file %v\n", fn)
		}
		for fr.Scan() {
			rawLine := fr.Text()

			// skip lines that only contain whitespace
			if strings.TrimSpace(rawLine) == "" {
				log.Debug("Ignoring whitespace line")
				continue
			}
			processLine(rawLine)
		}
		wg.Done()
	}()

}

func processLine(rawLine string) {
	log.Tracef("Working on line: %v\n", rawLine)

	// remove special chars
	for _, v := range specialChars {
		rawLine = strings.ReplaceAll(rawLine, v, "")
	}

	rawLine = strings.ToLower(rawLine) // cast to lowercacse

	line := strings.Split(rawLine, " ") // split into words

	log.Tracef("Formatted line: %v\n", line)

	mt.Lock()
	defer mt.Unlock()
	for _, v := range line {
		result[v]++
	}
}

func printResult() {
	fmt.Printf("%-10s%s\n", "Count", "Word")
	fmt.Printf("%-10s%s\n", "-----", "----")

	for w, c := range result {
		fmt.Printf("%-10v%s\n", c, w)
	}
}
