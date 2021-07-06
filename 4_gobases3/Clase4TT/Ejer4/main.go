package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(1000000)
	c1, c2, c3 := make(chan time.Duration), make(chan time.Duration), make(chan time.Duration)

	go insertSort(variable1, c1)
	go bubbleSort(variable1, c2)
	go selectSort(variable1, c3)
	fmt.Printf("100 items\n")
	fmt.Printf("Insert %d nanoseconds\nBubble %d nanoseconds\nselect %d nanoseconds\n\n", <-c1, <-c2, <-c3)

	go insertSort(variable2, c1)
	go bubbleSort(variable2, c2)
	go selectSort(variable2, c3)
	fmt.Printf("1000 items\n")
	fmt.Printf("Insert %d nanoseconds\nBubble %d nanoseconds\nselect %d nanoseconds\n\n", <-c1, <-c2, <-c3)

	go insertSort(variable3, c1)
	go bubbleSort(variable3, c2)
	go selectSort(variable3, c3)
	fmt.Printf("1000000 items\n")
	fmt.Printf("Insert %d nanoseconds\nBubble %d nanoseconds\nselect %d nanoseconds\n\n", <-c1, <-c2, <-c3)
}

func insertSort(arr []int, c chan time.Duration) {
	tmpArr := make([]int, len(arr))
	copy(tmpArr, arr)
	start := time.Now()

	var n = len(tmpArr)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if tmpArr[j-1] > tmpArr[j] {
				tmpArr[j-1], tmpArr[j] = tmpArr[j], tmpArr[j-1]
			}
			j = j - 1
		}
	}

	end := time.Now()
	c <- end.Sub(start)
}

func bubbleSort(arr []int, c chan time.Duration) {
	tmpArr := make([]int, len(arr))
	copy(tmpArr, arr)
	start := time.Now()

	for i := len(tmpArr); i > 0; i-- {
		for j := 1; j < i; j++ {
			if tmpArr[j-1] > tmpArr[j] {
				swap := tmpArr[j]
				tmpArr[j] = tmpArr[j-1]
				tmpArr[j-1] = swap
			}
		}
	}

	end := time.Now()
	c <- end.Sub(start)
}
func selectSort(arr []int, c chan time.Duration) {
	tmpArr := make([]int, len(arr))
	copy(tmpArr, arr)
	start := time.Now()

	var n = len(tmpArr)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if tmpArr[j] < tmpArr[minIdx] {
				minIdx = j
			}
		}
		tmpArr[i], tmpArr[minIdx] = tmpArr[minIdx], tmpArr[i]
	}

	end := time.Now()
	c <- end.Sub(start)
}

/*
Una go routine por cada ejecución de ordenamiento
Debo esperar a que terminen los ordenamientos de 100 números para seguir el de 1.000 y después el de 1.000.000.
Por último debo medir el tiempo de cada uno y mostrar en pantalla el resultado, para saber qué ordenamiento fue mejor para cada arreglo
*/
