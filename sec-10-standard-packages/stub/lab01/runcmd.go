package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type WriteCounter struct {
	writeCount int
	byteCount  int
}

func main() {

	wc := new(WriteCounter)

	for {
		// get command
		cmd := readCmd()

		// redirect stdin, stdout, and stderr of command
		cmd.Stdin = os.Stdin
		cmd.Stdout = wc
		cmd.Stderr = wc

		// run command
		cmd.Run()

		fmt.Println(wc)
	}

}

func readCmd() *exec.Cmd {

	for {
		fmt.Print(" #>  ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		input := strings.Fields(scanner.Text())
		if len(input) == 0 {
			continue
		}

		return exec.Command(input[0], input[1:]...)
	}
}

func (wc *WriteCounter) Write(b []byte) (n int, err error) {
	if wc == nil {
		return len(b), nil
	}

	wc.byteCount += len(b)
	wc.writeCount++

	return os.Stdout.Write(b)
}

func (wc *WriteCounter) String() string {
	return fmt.Sprintf("Write count: %v, bytes written: %v", wc.writeCount, wc.byteCount)
}
