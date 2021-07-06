package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result struct {
	listInt     []int
	workingTime time.Duration
}

func bubbleSort(list []int, c chan Result) {
	init := time.Now()
	len := len(list)
	for i := 0; i < len; i++ {
		for j := len - 1; j > i; j-- {
			if list[j] < list[j-1] {
				list[j], list[j-1] = list[j-1], list[j]
			}
		}
	}
	end := time.Now()
	workingTime := end.Sub(init)
	result := Result{listInt: list, workingTime: workingTime}
	c <- result
}

func insertionSort(list []int, c chan Result) {
	init := time.Now()
	for i := 1; i < len(list); i++ {
		j := i
		for j > 0 && list[j-1] > list[j] {
			if list[j-1] > list[j] {
				list[j-1], list[j] = list[j], list[j-1]
			}
			j--
		}
	}
	end := time.Now()
	workingTime := end.Sub(init)
	result := Result{listInt: list, workingTime: workingTime}
	c <- result
}

func selectionSort(list []int, c chan Result) {
	init := time.Now()
	for i := 0; i < len(list); i++ {
		min, index := list[i], i
		value1 := list[i]
		for j := i + 1; j < len(list); j++ {
			value2 := list[j]
			if value2 < min {
				min, index = value2, j
			}
		}
		if min != value1 {
			list[i], list[index] = min, value1
		}
	}
	end := time.Now()
	workingTime := end.Sub(init)
	result := Result{listInt: list, workingTime: workingTime}
	c <- result
}

func main() {
	variable1 := rand.Perm(100)
	//variable2 := rand.Perm(1000)
	//variable3 := rand.Perm(1000000)

	channel1 := make(chan Result)
	channel2 := make(chan Result)
	channel3 := make(chan Result)

	go bubbleSort(variable1, channel1)
	go insertionSort(variable1, channel2)
	go selectionSort(variable1, channel3)

	ch1 := <-channel1
	ch2 := <-channel2
	ch3 := <-channel3

	fmt.Println("Ordenamiento Burbuja: \n", ch1.listInt)
	fmt.Println("Tiempo Ordenamiento Burbuja: \n", ch1.workingTime)
	fmt.Println("Ordenamiento Inserccion: \n", ch2.listInt)
	fmt.Println("Tiempo Ordenamiento Inserccion: \n", ch2.workingTime)
	fmt.Println("Ordenamiento Seleccion: \n", ch3.listInt)
	fmt.Println("Tiempo Ordenamiento Seleccion: \n", ch3.workingTime)
}
