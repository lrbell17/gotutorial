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
	result       = make(map[string]int)
	specialChars = []string{"\"", ",", ".", "(", ")", "!", "?", "-", "]", "[", "'", ";", "*", "_", "=", ":", "+"}
	wg           sync.WaitGroup
)

func main() {

	if len(os.Args) == 1 {
		log.Fatal("No files to process")
	}

	inputFiles := os.Args[1:]

	start := time.Now()

	ch := inputReader(inputFiles)
	ch = lineFormatter(ch)
	ch = lineSplitter(ch)
	wordCounter(ch)

	wg.Wait()

	elapse := time.Since(start)
	printResult()
	fmt.Printf("Total time: %v\n", elapse)
}

// Generator to reads input file and publish lines to channel
func inputReader(inputFiles []string) (out chan string) {
	wg.Add(1)
	log.Infof("Processing %v files for input", len(inputFiles))

	out = make(chan string)

	go func() {
		defer wg.Done()
		defer close(out)

		for _, fn := range inputFiles {
			log.Infof("Processing file %v", fn)

			fr, frErr := input.NewFileReader(fn)

			// If a file can't be read, log error and continue to next file
			if frErr != nil {
				log.Errorf("Could not read file %v\n", fn)
				continue
			}

			for fr.Scan() {
				// fmt.Printf("Reading line: %v\n", fr.Text())
				out <- fr.Text()
			}

		}
	}()
	return
}

// Processor to format line and remove unwanted chars
func lineFormatter(in chan string) (out chan string) {

	wg.Add(1)
	out = make(chan string)

	go func() {
		defer wg.Done()
		defer close(out)

		for line := range in {
			if strings.TrimSpace(line) != "" { // ignore empty lines

				// remove special chars
				for _, v := range specialChars {
					line = strings.ReplaceAll(line, v, "")
				}

				// cast to lowercase & publish to output channel
				out <- strings.ToLower(line)

			}
		}
	}()

	return
}

// Processor to split lines into individual words
func lineSplitter(in chan string) (out chan string) {
	wg.Add(1)
	out = make(chan string)

	go func() {
		defer wg.Done()
		defer close(out)

		for line := range in {
			for _, word := range strings.Split(line, " ") {
				out <- word
			}
		}
	}()

	return
}

// Sink to count words
func wordCounter(in chan string) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		for word := range in {
			result[word]++
		}

	}()
}

func printResult() {
	fmt.Printf("%-10s%s\n", "Count", "Word")
	fmt.Printf("%-10s%s\n", "-----", "----")

	for w, c := range result {
		fmt.Printf("%-10v%s\n", c, w)
	}
}
