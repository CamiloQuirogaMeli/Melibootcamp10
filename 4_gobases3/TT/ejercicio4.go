package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)

	var c1 = make(chan []int)
	var c2 = make(chan []int)
	var c3 = make(chan []int)

	go InsertionSort(variable1, c1)
	go BubbleSort(variable1, c2)
	go SelectionSort(variable1, c3)

	fmt.Println("resultado insert 100 elementos:", <-c1)
	fmt.Println("resultado burbuja 100 elementos:", <-c2)
	fmt.Println("resultado seleccion 100 elementos:", <-c3)

	fmt.Println()

	go InsertionSort(variable2, c1)
	go BubbleSort(variable2, c2)
	go SelectionSort(variable2, c3)

	fmt.Println("resultado insert 1000 elementos:", <-c1)
	fmt.Println("resultado burbuja 1000 elementos:", <-c2)
	fmt.Println("resultado seleccion 1000 elementos:", <-c3)
}

func InsertionSort(list []int, c chan []int) {
	start := time.Now()
	stepCounter := 1
	for i, v := range list {
		if i < 1 {
			continue
		}
		for j := i - 1; j >= 0; j-- {
			stepCounter++
			if list[j] <= v {
				break
			}
			list[i] = list[j]
			list[j] = v
			i--
		}
	}
	fmt.Println("Termino Insert:", time.Since(start))
	c <- list
}

func BubbleSort(arr []int, c chan []int) {
	start := time.Now()
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("Termino Burbuja:", time.Since(start))
	c <- arr
}

func SelectionSort(arreglo []int, c chan []int) {
	start := time.Now()
	for i := 0; i < len(arreglo); i++ {
		minimo_encontrado, posicion_minimo := arreglo[i], i
		valor_original := arreglo[i]
		for j := i + 1; j < len(arreglo); j++ {
			valor_comparacion := arreglo[j]
			if valor_comparacion < minimo_encontrado {
				minimo_encontrado, posicion_minimo = valor_comparacion, j
			}
		}

		if minimo_encontrado != valor_original {
			arreglo[i], arreglo[posicion_minimo] = minimo_encontrado, valor_original
		}
	}
	fmt.Println("Termino Sleccion:", time.Since(start))
	c <- arreglo
}
