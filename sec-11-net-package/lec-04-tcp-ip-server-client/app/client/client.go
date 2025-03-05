package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"time"

	"github.com/lrbell17/gotutorial/sec-11-net-package/lec-04-tcp-ip-server-client/app/protocol"
	log "github.com/sirupsen/logrus"
)

var remoteAddr = "localhost:12345"

func main() {

	flag.StringVar(&remoteAddr, "s", remoteAddr, "random number server address")
	flag.Parse()

	consumer(remoteAddr)

}

func consumer(server string) {
	conn, err := net.Dial("tcp", server)
	if err != nil {
		log.Errorf("Can't connect to server %v: %v", server, err)
		return
	}

	jsonDec := json.NewDecoder(conn)
	var msg protocol.Message

	for {
		err = jsonDec.Decode(&msg) // decode message from json
		if err != nil {
			log.Errorf("Can't decode message for server %v: %v", server, err)
			return
		}
		processMessage(msg)
		conn.SetDeadline(time.Now().Add(5 * time.Second)) // reset timer after we receive work
	}

}

// calculate stats for message
func processMessage(msg protocol.Message) {
	total := 0

	for v := range msg.Payload {
		total += v
	}
	fmt.Printf("%v send:\n", msg.Source)
	l := len(msg.Payload)
	fmt.Printf("\tcount: %v, total: %v, avg: %v\n", l, total, float64(total)/float64(l))
}
