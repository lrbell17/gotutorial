package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	srvAddr1 = ":12345"
	srvAddr2 = ":12346"
)

type (
	// ServerOne and ServerTwo implements http.Handler
	ServerOne struct {
		counter int
	}
	ServerTwo struct {
		prefix string
	}
)

func main() {

	// fmt.Printf("Listening on %v\n", localhost)
	// http.HandleFunc("/", rootHandler)

	// http.Handle("/bar", new(ServerOne))
	// log.Fatal(http.ListenAndServe(localhost, new(ServerOne)))

	s1 := &http.Server{
		Addr:        srvAddr1,
		Handler:     new(ServerOne),
		IdleTimeout: 5 * time.Second,
	}
	s2 := &http.Server{
		Addr:        srvAddr2,
		Handler:     &ServerTwo{prefix: "Logger"},
		IdleTimeout: 10 * time.Second,
	}

	wg := new(sync.WaitGroup)
	startServer(wg, s1)
	startServer(wg, s2)
	wg.Wait()

}

func (s *ServerOne) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if s == nil {
		return
	}
	s.counter++
	fmt.Printf("ServerOne.ServeHTTP() called %v times\n", *s)
	fmt.Fprintf(w, "This is ServerOne being called %v times\n", *s)
}

func (s *ServerTwo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if s == nil {
		return
	}
	fmt.Fprintf(w, "[%s] - This is ServerTwo being called at %v", s.prefix, time.Now())
}

func startServer(wg *sync.WaitGroup, server *http.Server) {
	wg.Add(1)
	go func() {
		defer wg.Done()

		log.Infof("Starting server on address %v", server.Addr)
		err := server.ListenAndServe()
		if err != nil {
			log.Warnf("Server on address %v didn't start: %v", server.Addr, err)
			return
		}

	}()
}

// func rootHandler(w http.ResponseWriter, r *http.Request) {
// 	rootRequests++
// 	fmt.Fprintf(w, "This is rootHandler being called %v times", rootRequests)
// }
