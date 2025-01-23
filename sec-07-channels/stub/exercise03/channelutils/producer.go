package channelutils

import (
	"math/rand"
)

func LaunchProducers(ch chan Message, maxMessagePerProducer, numProducers int) {

	for i := 0; i < numProducers; i++ {
		producer(ch, id(i+1), maxMessagePerProducer)
	}
	close(ch)
}

func producer(ch chan Message, producerId id, maxMessagesPerProducer int) {

	numMessages := rand.Intn(maxMessagesPerProducer) + 1

	for i := 0; i < numMessages; i++ {
		num := rand.Intn(100) + 1
		ch <- Message{producerId, num}
	}
}
