package main

import (
	"encoding/json"
	"net"
	"os/exec"
	"time"

	"github.com/lrbell17/gotutorial/sec-11-net-package/stub/lab01/cmd"

	log "github.com/sirupsen/logrus"
)

const (
	localhost = ":12345"
)

func main() {
	ln, err := net.Listen("tcp", localhost)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	log.Infof("Command server listening %v", localhost)

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}
}
func handleConnection(conn net.Conn) {
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(5 * time.Second)) // wait 5 secs for command before closing connection

	// TODO 1 - complete this function
	// TODO 1.a read the command from the client in JSON, use local package cmd.Command and standard package *encoding/json*
	jsonEnc := json.NewDecoder(conn)
	var command cmd.Command
	decoderErr := jsonEnc.Decode(&command)
	if decoderErr != nil {
		log.Errorf("Can't decode command: %v", decoderErr)
	}

	// TODO 1.b execute the command on this host, where the server is running
	// TODO 1.c connect the command's stdin, stdout, and stderr to the connection
	cmd := exec.Command(command.Name, command.Args...)
	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn

	cmd.Run()

}
