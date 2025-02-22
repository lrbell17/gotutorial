package main

import (
	"fmt"
	"io"
	"os"
)

var data = []string{
	"Hello, good morning!",
	"Hi, my name is Sam",
	"Hey, good to see you again",
	"Welcome back",
}

type (
	ID     string
	Person struct {
		name string
	}
)

func main() {

	// write data to file
	file, _ := os.Create("data.txt")
	defer file.Close()

	for _, s := range data {
		io.WriteString(file, s)
		io.WriteString(file, "\n")
	}

	// write data to stdout
	out := os.Stdout
	defer out.Close()
	for _, s := range data {
		io.WriteString(out, s)
		io.WriteString(out, "\n")
	}

	// Println example
	Println(true)
	Println(8)
	Println(3.14)
	Println(4i + 1)
	Println(&Person{"John"})
	Println("hi")

	Fprintln(file, &Person{"John"})

}

func Println(v ...any) {
	Print(v...)
	Print("\n")
}

func Print(v ...any) {
	s := Sprint(v...)
	Fprint(os.Stdout, s)
}

func Fprintln(w io.Writer, v ...any) {
	Fprint(w, v...)
	Fprint(w, "\n")
}

func Fprint(w io.Writer, v ...any) {
	s := Sprint(v...)
	io.WriteString(w, s)
}

func Sprint(v ...any) (s string) {
	for i, x := range v {
		switch x := x.(type) { // checks type of x
		case bool:
			if x {
				s += "true"
			} else {
				s += "false"
			}
		case int:
			s += intToStr(x)
		case float64:
			s += f64ToStr(x)
		case complex128:
			s += c128ToStr(x)
		case *Person:
			s += ppToStr(x)
		case string:
			s += x
		default:
			s += unknownToStr(x)
		}
		if i > 0 {
			s += " "
		}
	}
	return
}

func intToStr(i int) string {
	return fmt.Sprint(i)
}
func f64ToStr(f float64) string {
	return fmt.Sprint(f)
}
func c128ToStr(c complex128) string {
	return fmt.Sprintf("(%v+%vi)", real(c), imag(c))
}
func ppToStr(p *Person) string {
	return fmt.Sprintf("%v", p)
}
func unknownToStr(v interface{}) string {
	return fmt.Sprintf("%#v.%T", v, v)

}
