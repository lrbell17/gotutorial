package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type (
	Message struct {
		Source  string
		Payload []int
	}
)

func main() {

	serverAddr := ":12345"
	remoteAddr := "localhost:12345"

	wg := new(sync.WaitGroup)
	producer(wg, "Producer 1", serverAddr)
	consumer(wg, remoteAddr)
	wg.Wait()
}

// consumer is a client that consumes messages from a server
func consumer(wg *sync.WaitGroup, server string) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		conn, err := net.Dial("tcp", server)
		if err != nil {
			log.Errorf("Can't connect to server %v: %v", server, err)
			return
		}

		jsonDec := json.NewDecoder(conn)
		var msg Message

		for {
			err = jsonDec.Decode(&msg) // decode message from json
			if err != nil {
				log.Errorf("Can't decode message for server %v: %v", server, err)
				return
			}
			processMessage(msg)
			conn.SetDeadline(time.Now().Add(5 * time.Second)) // reset timer after we receive work
		}
	}()

}

// producer is a server that provides messages
func producer(wg *sync.WaitGroup, src string, addr string) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		ln, err := net.Listen("tcp", addr)
		if err != nil {
			log.Errorf("Can't create server for %v: %v", addr, err)
		}
		defer ln.Close()
		log.Infof("Random number server started on %v", addr)

		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Errorf("Client connection error: %v", err)
				continue // get another connection to serve
			}
			go serveClient(conn, src) // create goroutine to serve client and keep accepting connections
		}
	}()
}

func serveClient(conn net.Conn, src string) {
	jsonEnc := json.NewEncoder(conn)
	var err error
	for {
		msg := newMessage(src)
		err = jsonEnc.Encode(msg) // encode message to json
		if err != nil {
			log.Errorf("Can't send to client %v, closing connection: %v", conn.RemoteAddr(), err)
			conn.Close()
			return
		}
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond) // random delay up to 2sec
		conn.SetDeadline(time.Now().Add(5 * time.Second))             // reset timer after sending
	}
}

// create random payload
func newMessage(src string) Message {
	count := rand.Intn(10) + 1 // 1 to 10
	data := make([]int, count)
	for i := 0; i < count; i++ {
		data[i] = rand.Intn(100)
	}
	return Message{Source: src, Payload: data}
}

// calculate stats for message
func processMessage(msg Message) {
	total := 0

	for v := range msg.Payload {
		total += v
	}
	fmt.Printf("%v send:\n", msg.Source)
	l := len(msg.Payload)
	fmt.Printf("\tcount: %v, total: %v, avg: %v\n", l, total, float64(total)/float64(l))
}
