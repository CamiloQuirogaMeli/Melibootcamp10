package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ordenamientoInserccion(arr []int, c chan time.Duration) {
	start := time.Now()
	var auxiliar int
	for i := 1; i < len(arr); i++ {
		auxiliar = arr[i]
		for j := i - 1; j >= 0 && arr[j] > auxiliar; j-- {
			arr[j+1] = arr[j]
			arr[j] = auxiliar
		}
	}
	duration := time.Since(start)
	c <- duration
}

func ordenamientoBurbuja(arr []int, c chan time.Duration) {
	start := time.Now()
	var auxiliar int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] > arr[j] {
				auxiliar = arr[i]
				arr[i] = arr[j]
				arr[j] = auxiliar
			}
		}
	}
	duration := time.Since(start)
	c <- duration
}

func intercambiar(a, b *int) {
	temporal := *a
	*a = *b
	*b = temporal
}

func ordenamientoSeleccion(arr []int, c chan time.Duration) {
	start := time.Now()
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {

			if arr[i] > arr[j] {
				intercambiar(&arr[i], &arr[j])
			}
		}
	}
	duration := time.Since(start)
	c <- duration
}

func main() {

	c := make(chan time.Duration)
	arr1, arr2, arr3 := rand.Perm(100), rand.Perm(1000), rand.Perm(1000000)

	fmt.Println("-> inicia ordenamiento INSERCION aray 1")
	go ordenamientoInserccion(arr1, c)
	fmt.Println("-> finaliza ordenamiento INSERCION aray 1, tardo ", <-c)
	fmt.Println("")
	fmt.Println("-> inicia ordenamiento BURBUJA aray 1")
	go ordenamientoBurbuja(arr1, c)
	fmt.Println("-> finaliza ordenamiento BURBUJA aray 1, tardo ", <-c)
	fmt.Println("")
	fmt.Println("-> inicia ordenamiento SELECCION aray 1")
	go ordenamientoSeleccion(arr1, c)
	fmt.Println("-> finaliza ordenamiento SELECCION aray 1, tardo ", <-c)

	fmt.Println("")

	fmt.Println("-> inicia ordenamiento INSERCION aray 2")
	go ordenamientoInserccion(arr2, c)
	fmt.Println("-> finaliza ordenamiento INSERCION aray 2, tardo ", <-c)
	fmt.Println("")
	fmt.Println("-> inicia ordenamiento BURBUJA aray 2")
	go ordenamientoBurbuja(arr2, c)
	fmt.Println("-> finaliza ordenamiento BURBUJA aray 2, tardo ", <-c)
	fmt.Println("")
	fmt.Println("-> inicia ordenamiento SELECCION aray 2")
	go ordenamientoSeleccion(arr2, c)
	fmt.Println("-> finaliza ordenamiento SELECCION aray 2, tardo ", <-c)

	fmt.Println("")

	fmt.Println("-> inicia ordenamiento INSERCION aray 3")
	go ordenamientoInserccion(arr3, c)
	fmt.Println("-> finaliza ordenamiento INSERCION aray 3, tardo ", <-c)
	fmt.Println("")
	fmt.Println("-> inicia ordenamiento BURBUJA aray 3")
	go ordenamientoBurbuja(arr3, c)
	fmt.Println("-> finaliza ordenamiento BURBUJA aray 3, tardo ", <-c)
	fmt.Println("")
	fmt.Println("-> inicia ordenamiento SELECCION aray 3")
	go ordenamientoSeleccion(arr3, c)
	fmt.Println("-> finaliza ordenamiento SELECCION aray 3, tardo ", <-c)

}
