package main

import (
	"fmt"
	"io"
	"sync"
)

var data = []string{"hello world!", "this is very nice", "how cool!"}

func main() {
	var wg sync.WaitGroup
	r, w := io.Pipe()
	producer(&wg, w)
	consumer(&wg, r)
	wg.Wait()
}

func consumer(wg *sync.WaitGroup, in io.Reader) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		s := make([]byte, 256)

		var err error
		var n int

		for err == nil {
			n, err = in.Read(s)

			if err == nil || (err == io.EOF && n > 0) { // if error == nil or EOF, print everything up to n
				fmt.Println(string(s[:n]))
			}

		}
	}()

}

func producer(wg *sync.WaitGroup, out io.WriteCloser) {
	wg.Add(1)
	go func() {
		defer out.Close() // close the writer
		defer wg.Done()

		for _, s := range data { // write data to channel
			io.WriteString(out, s)
		}
	}()
}
