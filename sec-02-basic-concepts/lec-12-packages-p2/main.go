package main

import (
	"fmt"
	// "./logger" --> relative import

	// alias - override package name
	log "github.com/lrbell17/gotutorial/sec-02-basic-concepts/lec-12-packages-p2/logger"

	// "." --> full import (not recommended)
	. "github.com/sirupsen/logrus"

	// "_" --> run init() but don't actually use the package
	_ "github.com/lrbell17/gotutorial/sec-02-basic-concepts/lec-12-packages-p2/math"
)

// can definie multiple init functions!
func init() {
	fmt.Println("main.init() from main package")
}

func init() {
	fmt.Println("main.init() from main package (part 2)")
}

func main() {
	log.Info("My info message")
	log.Error("My error message")

	Error("Logrus error message")

	// fmt.Println(math.Pi)
}
