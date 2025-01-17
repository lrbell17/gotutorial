package logger

import (
	"fmt"
)

func init() {
	fmt.Println("init() from logger package info.go")
}

func Info(s string) {
	fmt.Println("[INFO]", s)
}
