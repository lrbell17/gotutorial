package main

import (
	"flag"

	"github.com/lrbell17/gotutorial/sec-07-channels/stub/exercise03/channelutils"
)

var (
	maxMessagePerProducer = 100
	numProducers          = 3
)

func main() {

	if maxMessagePerProducer < 1 || numProducers < 1 {
		flag.Usage()
		return
	}

	ch := make(chan channelutils.Message, maxMessagePerProducer*numProducers)
	channelutils.LaunchProducers(ch, maxMessagePerProducer, numProducers)
	channelutils.Consumer(ch)

}

func init() {
	flag.IntVar(&maxMessagePerProducer, "m", maxMessagePerProducer, "The max messages per producer (>0)")
	flag.IntVar(&numProducers, "np", numProducers, "The number of producers (>0)")
	flag.Parse()
}
