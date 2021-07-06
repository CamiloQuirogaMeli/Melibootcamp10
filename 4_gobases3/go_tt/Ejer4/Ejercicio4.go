/*
Una empresa de sistemas requiere analizar qué algoritmos de ordenamiento utilizar para sus
servicios.
Para ellos se requiere instanciar 3 arreglos con valores aleatorios desordenados
- un arreglo de números enteros con 100 valores
- un arreglo de números enteros con 1.000 valores
- un arreglo de números enteros con 1.000.000 valores

Para instanciar las variables utilizar rand

Se debe realizar el ordenamiento de cada una por:
- Ordenamiento por inserción
- Ordenamiento por burbuja
- Ordenamiento por selección

Una go routine por cada ejecución de ordenamiento
Debo esperar a que terminen los ordenamientos de 100 números para seguir el de 1.000 y
después el de 1.000.000.
Por último debo medir el tiempo de cada uno y mostrar en pantalla el resultado, para saber
qué ordenamiento fue mejor para cada arreglo
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ordenamientoInsercion(arr []int, arrOrdenado chan []int) {
	start := time.Now()
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}
	stop := time.Now()
	tiempoEjecucion := stop.Sub(start)
	fmt.Println("El tiempo de ejecución del algoritmo 'Insercion' fue de: ", tiempoEjecucion)

	arrOrdenado <- arr
}

func ordenamientoBurbuja(arr []int, arrOrdenado chan []int) {
	start := time.Now()
	var flag bool
	for i := 0; i < len(arr)-1; i++ {
		flag = false
		for j := 0; j < len(arr)-1; j++ {
			if arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	stop := time.Now()
	tiempoEjecucion := stop.Sub(start)
	fmt.Println("El tiempo de ejecución del algoritmo 'Burbuja' fue de: ", tiempoEjecucion)
	arrOrdenado <- arr
}

func ordenamientoSeleccion(arr []int, arrOrdenado chan []int) {
	start := time.Now()
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
	stop := time.Now()
	tiempoEjecucion := stop.Sub(start)
	fmt.Println("El tiempo de ejecución del algoritmo 'Seleccion' fue de: ", tiempoEjecucion)
	arrOrdenado <- arr
}
func main() {
	arr1 := rand.Perm(100)
	arr2 := rand.Perm(1000)
	arr3 := rand.Perm(1000000)

	arrBurbujaOrdenado := make(chan []int)
	arrSeleccionOrdenado := make(chan []int)
	arrInsercionOrdenado := make(chan []int)

	for i := 0; i < 3; i++ {
		if i == 0 {
			fmt.Println("Incio array 1")
			go ordenamientoBurbuja(arr1, arrBurbujaOrdenado)
			go ordenamientoSeleccion(arr1, arrSeleccionOrdenado)
			go ordenamientoInsercion(arr1, arrInsercionOrdenado)
			fmt.Println("Fin array 1")
		} else if i == 1 {
			fmt.Println("Incio array 2")
			go ordenamientoBurbuja(arr2, arrBurbujaOrdenado)
			go ordenamientoSeleccion(arr2, arrSeleccionOrdenado)
			go ordenamientoInsercion(arr2, arrInsercionOrdenado)
			fmt.Println("Fin array 2")
		} else if i == 2 {
			fmt.Println("Incio array 3")
			go ordenamientoBurbuja(arr3, arrBurbujaOrdenado)
			go ordenamientoSeleccion(arr3, arrSeleccionOrdenado)
			go ordenamientoInsercion(arr3, arrInsercionOrdenado)
			fmt.Println("Fin array 3")
		}
		fmt.Println(<-arrBurbujaOrdenado)
		fmt.Println(<-arrSeleccionOrdenado)
		fmt.Println(<-arrInsercionOrdenado)
	}

	fmt.Println("Ejercicios finalizados")

}
