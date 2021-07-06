package main

import (
	"fmt"
	"math/rand"
	"time"
)

func insertionSort(array *[]int) {

	for i := 0; i < len(*array); i++ {

		for j := i; j > 0; j-- {

			i1 := (*array)[j-1]
			i2 := (*array)[j]

			if i2 < i1 {
				(*array)[j-1], (*array)[j] = i2, i1
			} else {
				break
			}

		}

	}

}

func selectionSort(array *[]int) {

	for i := 0; i < len(*array); i++ {

		minVal := (*array)[i]
		minInd := i

		for j := i; j < len(*array); j++ {

			val := (*array)[j]

			if val < minVal {

				minVal = val
				minInd = j

			}

		}
		(*array)[minInd], (*array)[i] = (*array)[i], (*array)[minInd]

	}

}

func bubbleSort(array *[]int) {

	changed := true

	for changed {

		changed = false

		for i := 1; i < len(*array); i++ {

			i1 := (*array)[i-1]
			i2 := (*array)[i]

			if i1 > i2 {

				(*array)[i-1], (*array)[i] = i2, i1
				changed = true

			}
		}
	}

}

func sortTimeRoutine(sortFunction func(array *[]int), arrayCount int, c chan time.Duration) {
	c <- sortTime(sortFunction, arrayCount)
}

func sortTime(sortFunction func(array *[]int), arrayCount int) time.Duration {

	array := rand.Perm(arrayCount)
	start := time.Now()

	sortFunction(&array)

	return time.Since(start)

}

func sortingTime(sizes []int) {
	for i := 0; i < 3; i++ {

		size := sizes[i]

		chanDurations := []chan time.Duration{make(chan time.Duration), make(chan time.Duration),
			make(chan time.Duration)}

		go sortTimeRoutine(bubbleSort, size, chanDurations[0])
		go sortTimeRoutine(insertionSort, size, chanDurations[1])
		go sortTimeRoutine(selectionSort, size, chanDurations[2])

		fmt.Println("\nSorting time for", size, "numbers: \n")
		fmt.Println("Bubble sort:", <-chanDurations[0])
		fmt.Println("Insertion sort:", <-chanDurations[1])
		fmt.Println("Selection sort:", <-chanDurations[2])
	}
}
