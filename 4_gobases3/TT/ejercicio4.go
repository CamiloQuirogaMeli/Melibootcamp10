package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//start := time.Now()
	//time.Sleep(5 )
	//end := time.Since(start)
	//fmt.Println(end )
	//os.Exit(1)

	rand1 := rand.Perm(100)
	rand2 := rand.Perm(1000)
	rand3 := rand.Perm(10000)

	s1 := "Slice de 100:"
	s2 := "Slice de 1000:"
	s3 := "Slice de 10000:"

	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go bubbleSort(rand1, c1)
	go insertionSort(rand1, c2)
	go selectionSort(rand1, c3)
	fmt.Println(s1, <-c1)
	fmt.Println(s1, <-c2)
	fmt.Println(s1, <-c3)

	go bubbleSort(rand2, c1)
	go insertionSort(rand2, c2)
	go selectionSort(rand2, c3)
	fmt.Println(s2, <-c1)
	fmt.Println(s2, <-c2)
	fmt.Println(s2, <-c3)

	go bubbleSort(rand3, c1)
	go insertionSort(rand3, c2)
	go selectionSort(rand3, c3)
	fmt.Println(s3, <-c1)
	fmt.Println(s3, <-c2)
	fmt.Println(s3, <-c3)

}

func bubbleSort(values []int, channel chan string) {
	start := time.Now()

	for i := len(values); i > 0; i-- {
		for j := 1; j < i; j++ {
			if values[j-1] > values[j] {
				intermediate := values[j]
				values[j] = values[j-1]
				values[j-1] = intermediate
			}
		}
	}
	duration := time.Since(start)
	channel <- fmt.Sprint("Bubble Sort -> ", duration)
}

func insertionSort(values []int, channel chan string) {
	start := time.Now()

	var n = len(values)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if values[j-1] > values[j] {
				values[j-1], values[j] = values[j], values[j-1]
			}
			j = j - 1
		}
	}
	duration := time.Since(start)
	channel <- fmt.Sprint("InsertionSort -> ", duration)
}

func selectionSort(values []int, channel chan string) {
	start := time.Now()

	var n = len(values)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if values[j] < values[minIdx] {
				minIdx = j
			}
		}
		values[i], values[minIdx] = values[minIdx], values[i]
	}

	duration := time.Since(start)
	channel <- fmt.Sprint("SelectionSort -> ", duration)
}