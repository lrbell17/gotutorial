package main

import (
	"encoding/json"
	"math/rand"
	"net"
	"sync"
	"time"

	"github.com/lrbell17/gotutorial/sec-11-net-package/lec-04-tcp-ip-server-client/app/protocol"
	log "github.com/sirupsen/logrus"
)

func main() {
	serverAddr := ":12345"
	wg := new(sync.WaitGroup)
	producer(wg, "Producer 1", serverAddr)
	wg.Wait()
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
func newMessage(src string) protocol.Message {
	count := rand.Intn(10) + 1 // 1 to 10
	data := make([]int, count)
	for i := 0; i < count; i++ {
		data[i] = rand.Intn(100)
	}
	return protocol.Message{Source: src, Payload: data}
}
