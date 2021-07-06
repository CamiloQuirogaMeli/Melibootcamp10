package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	timeNumbers := []string{"One Hundred", "One Thousand", "One Million"}
	numbers := make([][]int, 0)

	numbers = append(numbers, rand.Perm(100))
	numbers = append(numbers, rand.Perm(1000))
	numbers = append(numbers, rand.Perm(100000))

	ci := make(chan int)
	cb := make(chan int)
	cs := make(chan int)

	for i := 0; i < 3; i++ {
		fmt.Printf("Time for %s\n", timeNumbers[i])

		go insertionSort(numbers[i], ci)
		fmt.Println("Time InsertionSort ", <-ci)

		go bubbleSort(numbers[i], cb)
		fmt.Println("Time BubbleSort ", <-cb)

		go selectionSort(numbers[i], cs)
		fmt.Println("Time SelectionSort ", <-cs)
	}

}

func insertionSort(arr []int, c chan int) {
	initialize := time.Now()

	len := len(arr)
	for i := 1; i < len; i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}

	finalize := time.Since(initialize)
	c <- int(finalize)
}

func bubbleSort(arr []int, c chan int) {
	initialize := time.Now()

	len := len(arr)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	finalize := time.Since(initialize)
	c <- int(finalize)
}

func selectionSort(arr []int, c chan int) {
	initialize := time.Now()

	len := len(arr)
	for i := 0; i < len-1; i++ {
		minIndex := i
		for j := i + 1; j < len; j++ {
			if arr[j] < arr[minIndex] {
				arr[j], arr[minIndex] = arr[minIndex], arr[j]
			}
		}
	}

	finalize := time.Since(initialize)
	c <- int(finalize)
}
