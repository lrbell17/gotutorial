package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/lrbell17/gotutorial/sec-10-standard-packages/stub/exercise02/rw"
)

func main() {
	var file *os.File
	fmt.Printf("File value: %v, type: %T\n", file, file)

	// create a file
	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	// write data using file
	f.WriteString("hello\n")

	// write data using io
	io.WriteString(f, "world\n")

	// write data using Writer interfaces
	w, _ := rw.NewRecordWriter(f)
	io.WriteString(w, "hello again")

	fmt.Printf("Wrote data to file: %v\n", f.Name())

	f.Close()

	// Opening a file
	fn := "data.txt"
	if len(os.Args) > 1 {
		fn = os.Args[1]
	}

	f2, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	// reading bytes from a file
	// b := make([]byte, 500)
	// n, err := f2.Read(b)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(b[:n]))

	// We can also copy to a writer without having to specify the # of bytes
	io.Copy(os.Stdout, f2)

}
