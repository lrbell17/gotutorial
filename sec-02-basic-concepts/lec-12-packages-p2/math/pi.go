package math

import (
	"math"

	log "github.com/lrbell17/gotutorial/sec-02-basic-concepts/lec-12-packages-p2/logger"
)

var Pi = math.Pi

func init() {
	log.Info("init() from math package pi.go")
}
