package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

var (
	hostname string
	help     bool
)

type readCounter struct {
	counts map[string]int
}

func init() {
	flag.BoolVar(&help, "h", false, "show usage")
	flag.StringVar(&hostname, "s", "http://localhost:12345", "website to fetch, eg: https://google.com")
	flag.Parse()
}

func main() {
	if help {
		flag.Usage()
		os.Exit(0)
	}

	// TODO 1 - complete the main() function to fetch data from the server specified by 'hostname'
	response, err := http.Get(hostname)
	if err != nil {
		log.Fatal("Request failed!\n", err)
	}
	defer response.Body.Close()

	buf, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		log.Fatal("Error reading response: ", err)
	}
	pageStats(buf)

}

func pageStats(buf []byte) {

	// TODO 2 - Fetch the page and determine how many types of HTML tags are used alone with the page size
	fmt.Printf("Stats for Page from '%v' has:\n", hostname)

	rc := newReadCounter()
	rc.Read(buf)

	for element, count := range rc.counts {
		fmt.Printf("%v: %v\n", element, count)
	}
}

func newReadCounter() (rc *readCounter) {
	rc = new(readCounter)
	rc.counts = make(map[string]int)
	return
}

func (rc *readCounter) Read(b []byte) (n int, err error) {
	if rc == nil {
		return 0, errors.New("nil type")
	}

	tagRegex := regexp.MustCompile(`<([a-zA-Z][a-zA-Z0-9]*)\b[^>]*>`)
	matches := tagRegex.FindAllSubmatch(b, -1)

	for _, match := range matches {
		if len(match) > 1 {
			rc.counts[string(match[1])]++
		}
	}

	return len(b), nil
}
