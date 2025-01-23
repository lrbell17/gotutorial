package channelutils

import (
	"fmt"
	"sync"
)

type stats struct {
	count int
	sum   int
}

var (
	WgConsumers   sync.WaitGroup
	consumerMutex sync.Mutex
)

func LaunchConsumers(ch chan Message, numConsumers int) {

	for i := range numConsumers {
		consumer(ch, id(i+1))
	}
}

func consumer(ch chan Message, consumerId id) {

	WgConsumers.Add(1)

	go func() {
		producerStatMap := make(map[id]stats)

		for v := range ch {
			stat := producerStatMap[v.producerId]
			stat.count++
			stat.sum += v.num
			producerStatMap[v.producerId] = stat
		}

		printResult(producerStatMap, consumerId)

		WgConsumers.Done()
	}()

}

func printResult(producerStatMap map[id]stats, consumerId id) {

	var totalCount, totalSum int

	consumerMutex.Lock()
	fmt.Printf("--- Consumer #%v ---\n", consumerId)
	for id, stat := range producerStatMap {
		fmt.Printf("Producer #%v\n", id)
		fmt.Printf("\tNumber of elements: %v\n", stat.count)
		fmt.Printf("\tSub-total: %v\n", stat.sum)

		totalCount += stat.count
		totalSum += stat.sum
	}

	fmt.Printf("Total count: %v\n", totalCount)
	fmt.Printf("Total sum: %v\n\n", totalSum)
	consumerMutex.Unlock()
}
