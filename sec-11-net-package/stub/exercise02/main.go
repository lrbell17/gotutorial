package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type (
	HomePage struct {
		buf []byte
	}
)

const (
	localhost = ":12345"
	respFile  = "./index.html"
)

func main() {
	data, err := ioutil.ReadFile(respFile)
	if err != nil {
		log.Fatalf("Unable to read '%v': %v", respFile, err)
	}

	fmt.Printf("Server listening on %v\n", localhost)
	http.Handle("/", &HomePage{buf: data})
	http.HandleFunc("/greet", greetHandler)
	log.Fatal(http.ListenAndServe(localhost, nil))
}
func (hp *HomePage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if hp == nil {
		return
	}
	w.Write(hp.buf)
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><head><title>Good day!</title></head><body>")
	name := r.FormValue("name")
	// TODO 1 - say hi to the user and the current server date and time
	// <your code here>
	t := time.Now()
	fmt.Fprintf(w, "<p>Hello %v</p>", name)
	fmt.Fprintf(w, "<p>The date is %v and the time is %v</p>", t.Format(time.RFC1123)[:16], t.Format(time.Kitchen))

	// ---- NO changes below here
	fmt.Fprint(w, `<br/><a href="/">Home</a></body></html>`)
}
