package channelutils

import (
	"math/rand"
	"sync"
)

var WgProducers sync.WaitGroup

func LaunchProducers(ch chan Message, maxMessagePerProducer, numProducers int) {

	for i := 0; i < numProducers; i++ {
		producer(ch, id(i+1), maxMessagePerProducer)
	}
}

func producer(ch chan Message, producerId id, maxMessagesPerProducer int) {

	WgProducers.Add(1)
	go func() {
		numMessages := rand.Intn(maxMessagesPerProducer) + 1

		for i := 0; i < numMessages; i++ {
			num := rand.Intn(100) + 1
			ch <- Message{producerId, num}
		}
		WgProducers.Done()
	}()
}
