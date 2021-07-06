package main

import (
	"fmt"
	"math/rand"
)

func ordenamientoBurbuja(arreglo []int, c chan []int) {
	len := len(arreglo)
	for i := 0; i < len; i++ {
		for j := len - 1; j > i; j-- {
			if arreglo[j] < arreglo[j-1] {
				arreglo[j], arreglo[j-1] = arreglo[j-1], arreglo[j]
			}
		}
	}
	c <- arreglo
}
func ordenamientoInsercion(arreglo []int, c chan []int) {
	for i := 1; i < len(arreglo); i++ {
		j := i
		for j > 0 && arreglo[j-1] > arreglo[j] {
			intercambio(j-1, j, &arreglo)
			j--
		}
	}
	c <- arreglo
}

func intercambio(previo, actual int, puntero_arreglo *[]int) {
	arreglo := *puntero_arreglo
	copia := arreglo[actual]
	arreglo[actual] = arreglo[previo]
	arreglo[previo] = copia
}

func ordenamientoSeleccion(arreglo []int, c chan []int) {
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
	c <- arreglo
}

func main() {
	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(1000000)

	channel1 := make(chan []int)
	channel2 := make(chan []int)
	channel3 := make(chan []int)

	go ordenamientoBurbuja(variable1, channel1)
	go ordenamientoInsercion(variable2, channel2)
	go ordenamientoSeleccion(variable3, channel3)

	fmt.Println("Ordenamiento Burbuja: \n", <-channel1)
	fmt.Println("Ordenamiento Insercion: \n", <-channel2)
	fmt.Println("Ordenamiento selecion\n", <-channel3)

}
