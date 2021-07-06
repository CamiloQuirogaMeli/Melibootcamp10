package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ordenamientoPorInsercion(arr []int, c chan time.Duration) {
	t1 := time.Now()
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}
	t2 := time.Now()
	diff := t2.Sub(t1)
	c <- diff
}

func ordenamientoPorBurbuja(arr []int, c chan time.Duration) {
	t1 := time.Now()
	len := len(arr)
	for i := 0; i < len; i++ {
		for j := len - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	t2 := time.Now()
	diff := t2.Sub(t1)
	c <- diff
}

func ordenamientoPorSeleccion(arr []int, c chan time.Duration) {
	t1 := time.Now()
	var minIndex int
	for i := 0; i < len(arr)-1; i++ {
		minIndex = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
	t2 := time.Now()
	diff := t2.Sub(t1)
	c <- diff
}

func main() {
	c1 := make(chan time.Duration)
	c2 := make(chan time.Duration)
	c3 := make(chan time.Duration)

	variable1 := rand.Perm(100)

	go ordenamientoPorBurbuja(variable1, c1)
	go ordenamientoPorInsercion(variable1, c2)
	go ordenamientoPorSeleccion(variable1, c3)

	fmt.Println("Ordenamiento Burbuja 100", <-c1)
	fmt.Println("Ordenamiento Inserccion 100", <-c2)
	fmt.Println("Ordenamiento por Seleccion 100", <-c3)

	c4 := make(chan time.Duration)
	c5 := make(chan time.Duration)
	c6 := make(chan time.Duration)

	variable2 := rand.Perm(1000)

	go ordenamientoPorBurbuja(variable2, c4)
	go ordenamientoPorInsercion(variable2, c5)
	go ordenamientoPorSeleccion(variable2, c6)

	fmt.Println("Ordenamiento Burbuja 1000", <-c4)
	fmt.Println("Ordenamiento Inserccion 1000", <-c5)
	fmt.Println("Ordenamiento por Seleccion 1000", <-c6)

	c7 := make(chan time.Duration)
	c8 := make(chan time.Duration)
	c9 := make(chan time.Duration)

	variable3 := rand.Perm(1000000)

	go ordenamientoPorBurbuja(variable3, c7)
	go ordenamientoPorInsercion(variable3, c8)
	go ordenamientoPorSeleccion(variable3, c9)

	fmt.Println("Ordenamiento Burbuja 1000000", <-c7)
	fmt.Println("Ordenamiento Inserccion 1000000", <-c8)
	fmt.Println("Ordenamiento por Seleccion 1000000", <-c9)
}
