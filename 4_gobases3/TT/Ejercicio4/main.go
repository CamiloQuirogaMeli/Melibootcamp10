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

	channel1 := make(chan []int)
	channel2 := make(chan []int)
	channel3 := make(chan []int)

	// ORDENAMIENTO BURBUJA
	fmt.Println("=======BURBUJA=======")
	tiempoInicio := time.Now()
	go ordenamientoBurbuja(variable1, channel1)
	tiempoFin := time.Now()
	demora := tiempoFin.Sub(tiempoInicio)
	fmt.Println("-- Tiempo ejecucion1: ", demora, " - Ordenamiento Burbuja1: ", <-channel1)

	go ordenamientoBurbuja(variable2, channel2)
	tiempoFin2 := time.Now()
	demora2 := tiempoFin2.Sub(tiempoInicio)
	fmt.Println("-- Tiempo ejecucion2: ", demora2, " - Ordenamiento Burbuja2: ", <-channel2)
	go ordenamientoBurbuja(variable3, channel3)
	tiempoFin3 := time.Now()
	demora3 := tiempoFin3.Sub(tiempoInicio)
	fmt.Println("-- Tiempo ejecucion3: ", demora3, " - Ordenamiento Burbuja3: ", <-channel3)

	fmt.Println("=======INSERCION=======")
	//	go ordenamientoInsercion(variable2, channel2)
	//	go ordenamientoSeleccion(variable3, channel3)

	//	fmt.Println("Burbuja:", <-channel1)
	//	fmt.Println("Insercion:", <-channel2)
	//	fmt.Println("Ordenamiento:", <-channel3)
}

func ordenamientoBurbuja(numeros []int, c chan []int) {
	for i := 0; i < len(numeros); i++ {
		for j := len(numeros) - 1; j > i; j-- {
			if numeros[j] < numeros[j-1] {
				numeros[j], numeros[j-1] = numeros[j-1], numeros[j]
			}
		}
	}
	c <- numeros
}

func ordenamientoInsercion(numeros []int, c chan []int) {
	for i := 1; i < len(numeros); i++ {
		j := i
		for j > 0 && numeros[j-1] > numeros[j] {
			intercambiar(j-1, j, &numeros)
			j--
		}
	}
	c <- numeros
}

func intercambiar(anterior, actual int, puntero *[]int) {
	numeros := *puntero
	copia := numeros[actual]
	numeros[actual] = numeros[anterior]
	numeros[anterior] = copia
}

func ordenamientoSeleccion(numeros []int, c chan []int) {
	for i := 0; i < len(numeros); i++ {
		minimo_encontrado, posicion_minimo := numeros[i], i

		valor_original := numeros[i]
		for j := i + 1; j < len(numeros); j++ {
			valor_comparacion := numeros[j]
			if valor_comparacion < minimo_encontrado {
				minimo_encontrado, posicion_minimo = valor_comparacion, j
			}
		}

		if minimo_encontrado != valor_original {
			numeros[i], numeros[posicion_minimo] = minimo_encontrado, valor_original
		}
	}
	c <- numeros
}
