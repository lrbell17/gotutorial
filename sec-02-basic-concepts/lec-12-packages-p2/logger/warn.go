package logger

import (
	"fmt"
)

func init() {
	fmt.Println("init() from logger package warn.go")
}

func Warn(s string) {
	fmt.Println("[WARN]", s)
}
