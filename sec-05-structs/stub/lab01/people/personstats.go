package people

import (
	"fmt"
	"math/rand"
)

func (people People) GetStatsForGender(gender string) {

	filteredPeople := FilterByGender(people, gender)
	filteredPeople.QuickSort(0, len(filteredPeople)-1)

	fmt.Printf("\n----- Salary stats for %v workers -----\n", gender)
	fmt.Printf("Total workers: %v\n", len(filteredPeople))
	fmt.Printf("Minimum salary: %v\n", filteredPeople[0].salary)
	fmt.Printf("Maximum salary: %v\n", filteredPeople[len(filteredPeople)-1].salary)
	fmt.Printf("Average salary: %v\n", filteredPeople.GetAvgSalary())

}

func (ppl People) GetAvgSalary() (salary Curreny) {

	var sum Curreny
	for _, person := range ppl {
		sum += person.salary
	}
	salary = sum / Curreny(len(ppl))
	return
}

func FilterByGender(ppl People, gender string) (filteredPeople People) {

	for _, person := range ppl {
		if person.gender == gender {
			filteredPeople = append(filteredPeople, person)
		}
	}
	return
}

func (people People) BubbleSort() {

	for i := 0; i < len(people)-1; i++ {
		for j := i; j < len(people); j++ {
			if people[i].salary > people[j].salary {
				people[i], people[j] = people[j], people[i]
			}
		}
	}
}

func (ppl People) QuickSort(low, high int) {
	if low < high {
		partitionIdx := ppl.partition(low, high)

		ppl.QuickSort(partitionIdx+1, high)
		ppl.QuickSort(low, partitionIdx)
	}

}

func (ppl People) partition(low, high int) int {

	pivotIdx := rand.Intn(high-low+1) + low
	ppl[pivotIdx], ppl[high] = ppl[high], ppl[pivotIdx]
	pivot := ppl[high].salary

	i := low

	for j := low; j < high; j++ {
		if ppl[j].salary < pivot {
			ppl[i], ppl[j] = ppl[j], ppl[i]
			i++
		}
	}

	ppl[high], ppl[i] = ppl[i], ppl[high]

	return i
}
