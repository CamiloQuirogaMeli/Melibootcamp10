package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var option int
	var aux int = 1

	rand1 := rand.Perm(100)
	rand2 := rand.Perm(1000)
	rand3 := rand.Perm(10000)

	for aux != 0 {

		const banner string = "\n███╗   ███╗    ███████╗    ███╗   ██╗    ██╗   ██╗\n████╗ ████║    ██╔════╝    ████╗  ██║    ██║   ██║\n██╔████╔██║    █████╗      ██╔██╗ ██║    ██║   ██║\n██║╚██╔╝██║    ██╔══╝      ██║╚██╗██║    ██║   ██║\n██║ ╚═╝ ██║    ███████╗    ██║ ╚████║    ╚██████╔╝\n╚═╝     ╚═╝    ╚══════╝    ╚═╝  ╚═══╝     ╚═════╝ "

		const options string = "1. Ejercutar ordenamiento \n0. Salir"

		fmt.Println(banner)
		fmt.Println(options)
		fmt.Scan(&option)

		switch option {
		case 1:

			c1 := make(chan string)
			c2 := make(chan string)
			c3 := make(chan string)
			go BubbleSort(rand1, c1)
			go insertionsort(rand1, c2)
			go selectionsort(rand1, c3)
			fmt.Println("100", <-c1)
			fmt.Println("100", <-c2)
			fmt.Println("100", <-c3)

			go BubbleSort(rand2, c1)
			go insertionsort(rand2, c2)
			go selectionsort(rand2, c3)
			fmt.Println("1000", <-c1)
			fmt.Println("1000", <-c2)
			fmt.Println("1000", <-c3)

			go BubbleSort(rand3, c1)
			go insertionsort(rand3, c2)
			go selectionsort(rand3, c3)
			fmt.Println("10000", <-c1)
			fmt.Println("10000", <-c2)
			fmt.Println("10000", <-c3)
		case 0:
			aux = 0
		default:
			fmt.Println("Opcion incorrecta!!!")
		}
	}

}

func BubbleSort(numbers []int, c chan string) {
	start := time.Now()
	// Code to measure

	//Start the loop in reverse order, so the loop will start with length
	//which is equal to the length of input array and then loop untill
	//reaches 1
	for i := len(numbers); i > 0; i-- {
		//The inner loop will first iterate through the full length
		//the next iteration will be through n-1
		// the next will be through n-2 and so on
		for j := 1; j < i; j++ {
			if numbers[j-1] > numbers[j] {
				intermediate := numbers[j]
				numbers[j] = numbers[j-1]
				numbers[j-1] = intermediate
			}
		}
	}
	duration := time.Since(start)
	c <- fmt.Sprintf("Total de tiempo Bubble Sort %d", duration)
}

func insertionsort(items []int, c chan string) {
	start := time.Now()

	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
	duration := time.Since(start)
	c <- fmt.Sprintf("Total de tiempo Insertion Sort %d", duration)
}

func selectionsort(items []int, c chan string) {
	start := time.Now()

	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}

	duration := time.Since(start)
	c <- fmt.Sprintf("Total de tiempo Selection Sort %d", duration)
}
