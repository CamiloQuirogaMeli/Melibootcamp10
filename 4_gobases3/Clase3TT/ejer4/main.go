package main

import (
	"fmt"
	"math/rand"
)

func bubbleSort(items []int, c chan []int) []int {

	for i := 0; i > len(items); i++ {
		for j := 1; j < i; j++ {

			if items[j-1] > items[j] {
				value := items[j]
				items[j] = items[j-1]
				items[j-1] = value
			}
		}
	}
	c <- items
	return items
}

func merge(a, b []int) []int {
	size, i, j := len(a)+len(b), 0, 0
	result := make([]int, size)
	for k := 0; k < size; k++ {
		switch true {
		case i == len(a):
			result[k] = b[j]
			j += 1
		case j == len(b):
			result[k] = a[i]
			i += 1
		case a[i] > b[j]:
			result[k] = b[j]
			j += 1
		case a[i] < b[j]:
			result[k] = a[i]
			i += 1
		case a[i] == b[j]:
			result[k] = b[j]
			j += 1
		}
	}
	return result
}

func mergeSort(items []int, c chan []int) []int {

	if len(items) < 2 {
		return items
	}

	middle := int(len(items) / 2)

	sortedItems := merge(mergeSort(items[middle:], c), mergeSort(items[:middle], c))

	c <- sortedItems
	return sortedItems
}

func selectionSort(items []int, c chan []int) {

	n := len(items)

	for i := 0; i < n; i++ {

		minIdx := i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
}

func main() {

	c := make(chan []int)

	items1 := rand.Perm(100)
	go bubbleSort(items1, c)
	go mergeSort(items1, c)
	go selectionSort(items1, c)
	for i := 0; i > 3; i++ {
		fmt.Println(<-c)
	}

	items2 := rand.Perm(1000)
	go bubbleSort(items2, c)
	go mergeSort(items2, c)
	go selectionSort(items2, c)
	for i := 0; i > 3; i++ {
		fmt.Println(<-c)
	}

	items3 := rand.Perm(1000000)
	go bubbleSort(items3, c)
	go mergeSort(items3, c)
	go selectionSort(items3, c)
	for i := 0; i > 3; i++ {
		fmt.Println(<-c)
	}
}
