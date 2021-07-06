package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

func insertionsort(items *[]int, c chan *[]int) {
	var n = len(*items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if (*items)[j-1] > (*items)[j] {
				(*items)[j-1], (*items)[j] = (*items)[j], (*items)[j-1]
			}
			j = j - 1
		}
	}
	defer wg.Done()
	c <- items
}

func bubblesort(items *[]int, c chan *[]int) {
	var (
		n      = len(*items)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if (*items)[i] > (*items)[i+1] {
				(*items)[i+1], (*items)[i] = (*items)[i], (*items)[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
	defer wg.Done()
	c <- items
}

func selectionsort(items *[]int, c chan *[]int) {
	var n = len(*items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if (*items)[j] < (*items)[minIdx] {
				minIdx = j
			}
		}
		(*items)[i], (*items)[minIdx] = (*items)[minIdx], (*items)[i]
	}
	defer wg.Done()
	c <- items
}

func main() {

	list1 := rand.Perm(10)
	list2 := rand.Perm(100)
	list3 := rand.Perm(1000)
	insertionChannel := make(chan *[]int, 3)
	wg.Add(3)
	go insertionsort(&list1, insertionChannel)
	go insertionsort(&list2, insertionChannel)
	go insertionsort(&list3, insertionChannel)
	wg.Wait()
	close(insertionChannel)
	fmt.Println("Insertion sort")
	for item := range insertionChannel {
		fmt.Println(*item)
	}
	fmt.Println("============================")
	fmt.Println()

	list4 := rand.Perm(10)
	list5 := rand.Perm(100)
	list6 := rand.Perm(1000)
	bubbleChannel := make(chan *[]int, 3)
	wg.Add(3)
	go bubblesort(&list4, bubbleChannel)
	go bubblesort(&list5, bubbleChannel)
	go bubblesort(&list6, bubbleChannel)
	wg.Wait()
	close(bubbleChannel)
	fmt.Println("Bubble sort")
	for item := range bubbleChannel {
		fmt.Println(*item)
	}
	fmt.Println("============================")
	fmt.Println()

	list7 := rand.Perm(10)
	list8 := rand.Perm(100)
	list9 := rand.Perm(1000)
	selectionChannel := make(chan *[]int, 3)
	wg.Add(3)
	go selectionsort(&list7, selectionChannel)
	go selectionsort(&list8, selectionChannel)
	go selectionsort(&list9, selectionChannel)
	wg.Wait()
	close(selectionChannel)
	fmt.Println("Selection sort")
	for item := range selectionChannel {
		fmt.Println(*item)
	}
	fmt.Println("============================")
	fmt.Println()

}
