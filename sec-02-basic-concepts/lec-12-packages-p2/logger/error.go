package logger

import (
	"fmt"
)

func init() {
	fmt.Println("init() from logger package error.go")
}

func Error(s string) {
	fmt.Println("[ERROR]", s)
}
