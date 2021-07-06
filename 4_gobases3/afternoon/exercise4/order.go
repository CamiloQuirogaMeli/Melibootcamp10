package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	BUBBLE    = "BUBBLE"
	INSERTION = "INSERTION"
	SELECTION = "SELECTION"
)

func main() {

	list := rand.Perm(100)
	fastestSorting := compareSortings(list)
	fmt.Println("For 100 elements fastest sorting algorithm is", fastestSorting)

	list = rand.Perm(1000)
	fastestSorting = compareSortings(list)
	fmt.Println("For 1000 elements fastest sorting algorithm is", fastestSorting)

	list = rand.Perm(1000000)
	fastestSorting = compareSortings(list)
	fmt.Println("For 1000000 elements fastest sorting algorithm is", fastestSorting)

}

func compareSortings(list []int) string {

	bubbleChannel := make(chan time.Duration)
	go messureSort(BUBBLE, list, bubbleChannel)

	insertionChannel := make(chan time.Duration)
	go messureSort(INSERTION, list, insertionChannel)

	selectionChannel := make(chan time.Duration)
	go messureSort(SELECTION, list, selectionChannel)

	bubbleDuration := <-bubbleChannel
	insertionDuration := <-insertionChannel
	selectionDuration := <-selectionChannel

	if bubbleDuration < insertionDuration && bubbleDuration < selectionDuration {
		return BUBBLE
	} else if insertionDuration < bubbleDuration && insertionDuration < selectionDuration {
		return INSERTION
	} else {
		return SELECTION
	}

}

func messureSort(sortType string, list []int, c chan time.Duration) {
	start := time.Now()
	list = sortFunction(BUBBLE)(list)
	duration := time.Since(start)
	fmt.Println(sortType, "sort", duration)
	c <- duration
}

func sortFunction(funcType string) func(list []int) []int {
	funcMap := map[string]func(list []int) []int{
		BUBBLE:    bubbleSort,
		INSERTION: insertionSort,
		SELECTION: selectionSort}

	return funcMap[funcType]
}

func bubbleSort(list []int) []int {
	n := len(list) - 1
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			list[j], list[j+1] = swap(list[j], list[j+1])
		}
	}
	return list
}

func insertionSort(list []int) []int {
	var n = len(list)
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			list[j-1], list[j] = swap(list[j-1], list[j])
		}
	}
	return list
}

func selectionSort(list []int) []int {
	var n = len(list)
	for i := 0; i < n; i++ {
		minId := i
		for j := i; j < n; j++ {
			if list[i] > list[j] {
				minId = j
			}
		}
		list[i], list[minId] = swap(list[i], list[minId])
	}
	return list
}

func swap(a int, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b

}
