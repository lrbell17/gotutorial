package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

const (
	localhost = ":12345"
)

var (
	rootRequests int
	fooRequests  int
	barRequests  int
)

func main() {
	// serving dynamic data
	fmt.Printf("Server listening on: %v\n", localhost)
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/foo", fooHandler)
	http.HandleFunc("/bar", barHandler)
	log.Fatal(http.ListenAndServe(localhost, nil))

}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	// get info from request
	remoteHost := r.RemoteAddr
	remoteUri := r.RequestURI
	method := r.Method

	// return status code
	w.WriteHeader(http.StatusForbidden)

	// write to responsewriter
	fmt.Fprintf(w, "<h1>This is rootHandler being called %v times</h1>", rootRequests)
	fmt.Fprintf(w, "<h2>Remote host/client addr: %v</h2>", remoteHost)
	fmt.Fprintf(w, "<h2>Remote host/client URI: %v</h2>", remoteUri)
	fmt.Fprintf(w, "<h2>Request made using HTTP method: %v</h2>", method)

	// read request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("No data from client")
		fmt.Fprintf(w, "<p>No data provided</p>")
	}
	r.Body.Close() // close body

	fmt.Printf("Client sent: %v\n", string(body))
	caps := strings.Title(string(body))
	fmt.Fprintf(w, "<p><b>Input: </b>%v</p>", string(body))
	fmt.Fprintf(w, "<hr/>")
	fmt.Fprintf(w, "<p><b>Output: </b>%v</p>", caps)
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	fooRequests++
	fmt.Fprintf(w, "This is fooHandler being called %v times", fooRequests)
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	barRequests++
	fmt.Fprintf(w, "This is barHandler being called %v times", barRequests)
}
