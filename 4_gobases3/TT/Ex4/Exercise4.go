package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sliceFactory() ([]int, []int, []int) {
	slice1 := rand.Perm(100)
	slice2 := rand.Perm(1000)
	slice3 := rand.Perm(1000000)

	return slice1, slice2, slice3
}

func insertionSort(data *[]int, c chan []int) {
	for i, v := range *data {
		j := i
		for ; j > 0 && v < (*data)[j-1]; j-- {
			(*data)[j] = (*data)[j-1]
		}
		(*data)[j] = v
	}
	c <- (*data)
}

func bubbleSort(data *[]int, c chan []int) {
	for i := len((*data)); i > 0; i-- {
		for j := 1; j < i; j++ {
			if (*data)[j-1] > (*data)[j] {
				intermediate := (*data)[j]
				(*data)[j] = (*data)[j-1]
				(*data)[j-1] = intermediate
			}
		}
	}

	c <- (*data)
}

func selectionSort(data *[]int, c chan []int) {
	var n = len(*data)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if (*data)[j] < (*data)[minIdx] {
				minIdx = j
			}
		}
		(*data)[i], (*data)[minIdx] = (*data)[minIdx], (*data)[i]
	}
	c <- (*data)
}

func main() {

	insChan := make(chan []int)
	bubChan := make(chan []int)
	selChan := make(chan []int)

	slice100, slice1000, slice1000000 := sliceFactory()
	slice100Ins, slice1000Ins, slice1000000Ins := append([]int(nil), slice100...), append([]int(nil), slice1000...), append([]int(nil), slice1000000...)
	slice100Bub, slice1000Bub, slice1000000Bub := append([]int(nil), slice100...), append([]int(nil), slice1000...), append([]int(nil), slice1000000...)
	slice100Sel, slice1000Sel, slice1000000Sel := append([]int(nil), slice100...), append([]int(nil), slice1000...), append([]int(nil), slice1000000...)

	fmt.Println("---------------------------")
	fmt.Println("Insertion Sort!")
	insStart := time.Now()

	go insertionSort(&slice100Ins, insChan)
	<-insChan
	go insertionSort(&slice1000Ins, insChan)
	<-insChan
	go insertionSort(&slice1000000Ins, insChan)
	<-insChan

	insElapsed := time.Since(insStart)
	fmt.Println("Time:", insElapsed)

	fmt.Println("---------------------------")
	fmt.Println("Bubble Sort!")
	bubStart := time.Now()

	go bubbleSort(&slice100Bub, bubChan)
	<-bubChan
	go bubbleSort(&slice1000Bub, bubChan)
	<-bubChan
	go bubbleSort(&slice1000000Bub, bubChan)
	<-bubChan

	bubElapsed := time.Since(bubStart)
	fmt.Println("Time:", bubElapsed)

	fmt.Println("---------------------------")
	fmt.Println("Selection Sort!")
	selStart := time.Now()

	go selectionSort(&slice100Sel, selChan)
	<-selChan
	go selectionSort(&slice1000Sel, selChan)
	<-selChan
	go selectionSort(&slice1000000Sel, selChan)
	<-selChan

	selElapsed := time.Since(selStart)
	fmt.Println("Time:", selElapsed)
}
