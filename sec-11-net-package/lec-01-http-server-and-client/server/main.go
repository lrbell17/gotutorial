package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const localhost = ":12345"

func main() {

	// create HTTP server
	fmt.Printf("Server listening on %v\n", localhost)
	http.HandleFunc("/", handler)

	// listen for incoming requests and serve them
	log.Fatal(http.ListenAndServe(localhost, nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	const respText = `
		<html>
			<p>Hello world</p>
		</html>
	`
	io.WriteString(w, respText)
}
