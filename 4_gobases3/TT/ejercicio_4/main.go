package main

import (
	"fmt"
	"math/rand"
	"time"
)
 
func MeasureSortTime(
	sortFunction func(list *[]int),
	listCount int,
) time.Duration {

	list := rand.Perm(listCount)
	start := time.Now()

	sortFunction(&list)

	return time.Since(start)
}

func MeasureSortTimeRoutine(
	sortFunction func(list *[]int),
	listCount int,
	ch chan time.Duration,
) {
	
	ch <- MeasureSortTime(sortFunction, listCount)
}

func main() {

	listSizes := []int{100, 1000, 1000000}
	
	for i := 0; i < 3; i++ {

		chDurations := []chan time.Duration{
			make(chan time.Duration),
			make(chan time.Duration),
			make(chan time.Duration),
		}
		
		size := listSizes[i]

		go MeasureSortTimeRoutine(BubbleSort, size, chDurations[0])
		go MeasureSortTimeRoutine(InsertionSort, size, chDurations[1])
		go MeasureSortTimeRoutine(SelectionSort, size, chDurations[2])

		fmt.Println("Sorting time for size:", size)
		fmt.Println("\tBubble sort:", <- chDurations[0])
		fmt.Println("\tInsertion sort:", <- chDurations[1])
		fmt.Println("\tSelection sort:", <- chDurations[2])
	}
}