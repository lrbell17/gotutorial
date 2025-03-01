package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	localhost = "http://localhost:12345"
	host      = "https://google123refwd.com"
)

func main() {

	// client - send GET request
	response, err := http.Get(host)
	if err != nil {
		log.Printf("Can't connect to %v: %v", host, err)
		log.Printf("Trying: %v", localhost)
		response, err = http.Get(localhost)
	}
	if err != nil {
		log.Fatal(err)
	}

	// close response
	defer response.Body.Close()

	io.Copy(os.Stdout, response.Body)
	fmt.Printf("Request completed with response code: %v\n", response.StatusCode)

}
