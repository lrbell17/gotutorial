package channelutils

import "fmt"

type stats struct {
	count int
	sum   int
}

func Consumer(ch chan Message) {

	producerStatMap := make(map[id]stats)

	for v := range ch {
		stat := producerStatMap[v.producerId]
		stat.count++
		stat.sum += v.num
		producerStatMap[v.producerId] = stat
	}

	printResult(producerStatMap)
}

func printResult(producerStatMap map[id]stats) {

	var totalCount, totalSum int

	for id, stat := range producerStatMap {
		fmt.Printf("Producer #%v\n", id)
		fmt.Printf("\tNumber of elements: %v\n", stat.count)
		fmt.Printf("\tSub-total: %v\n", stat.sum)

		totalCount += stat.count
		totalSum += stat.sum
	}

	fmt.Printf("Total count: %v\n", totalCount)
	fmt.Printf("Total sum: %v\n", totalSum)
}
