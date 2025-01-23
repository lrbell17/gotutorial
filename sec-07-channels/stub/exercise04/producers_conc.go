package main

import (
	"flag"

	"github.com/lrbell17/gotutorial/sec-07-channels/stub/exercise04/channelutils"
)

var (
	maxMessagePerProducer = 100
	numProducers          = 3
	numConsumers          = 2
)

func main() {

	if maxMessagePerProducer < 1 || numProducers < 1 || numConsumers < 1 {
		flag.Usage()
		return
	}

	ch := make(chan channelutils.Message) // unbuffered channel

	// Launch consumers and producers
	channelutils.LaunchConsumers(ch, numConsumers)
	channelutils.LaunchProducers(ch, maxMessagePerProducer, numProducers)

	channelutils.WgProducers.Wait() // wait for all producers
	close(ch)                       // close channel
	channelutils.WgConsumers.Wait() // wait for all cconsumers
}

func init() {
	flag.IntVar(&maxMessagePerProducer, "m", maxMessagePerProducer, "The max messages per producer (>0)")
	flag.IntVar(&numProducers, "np", numProducers, "The number of producers (>0)")
	flag.IntVar(&numConsumers, "nc", numConsumers, "The number of consumers (>0)")
	flag.Parse()
}
