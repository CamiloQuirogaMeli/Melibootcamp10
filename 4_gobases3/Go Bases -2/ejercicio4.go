package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	Ejercicio 4 - Ordenamiento
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

func main() {
	arreglo1 := rand.Perm(100)
	arreglo2 := rand.Perm(1000)
	arreglo3 := rand.Perm(1000000)

	canal1 := make(chan int)

	fmt.Println("Ordenamiento para array de 100 elementos")
	go ordenamientoInsercion(canal1, arreglo1)
	go ordenamientoBurbuja(canal1, arreglo1)
	go ordenamientoSeleccion(canal1, arreglo1)
	for i := 0; i < 3; i++ {
		<-canal1
	}

	fmt.Println("Ordenamiento para array de 1.000 elementos")
	go ordenamientoInsercion(canal1, arreglo2)
	go ordenamientoBurbuja(canal1, arreglo2)
	go ordenamientoSeleccion(canal1, arreglo2)
	for i := 0; i < 3; i++ {
		<-canal1
	}

	fmt.Println("Ordenamiento para array de 1'000.000 elementos")
	go ordenamientoInsercion(canal1, arreglo3)
	go ordenamientoBurbuja(canal1, arreglo3)
	go ordenamientoSeleccion(canal1, arreglo3)
	for i := 0; i < 3; i++ {
		<-canal1
	}
}

func ordenamientoInsercion(canal1 chan int, arreglo []int) []int {
	inicio := time.Now()
	//fmt.Println("El metodo inicio: ", inicio)
	for i := 0; i < len(arreglo); i++ {
		auxiliar := arreglo[i]
		for j := i - 1; j >= 0 && arreglo[j] > auxiliar; j-- {
			arreglo[j+1] = arreglo[j]
			arreglo[j] = auxiliar
		}
	}
	final := time.Since(inicio)
	fmt.Println("El metodo insercion tardo: ", final)
	canal1 <- 1
	//fmt.Println(arreglo)
	return arreglo
}

func ordenamientoBurbuja(canal1 chan int, arreglo []int) []int {
	var auxiliar int
	inicio := time.Now()
	//fmt.Println("El metodo inicio: ", inicio)
	for i := 0; i < len(arreglo); i++ {
		for j := 0; j < len(arreglo); j++ {
			if arreglo[i] > arreglo[j] {
				auxiliar = arreglo[i]
				arreglo[i] = arreglo[j]
				arreglo[j] = auxiliar
			}
		}
	}
	final := time.Since(inicio)
	fmt.Println("El metodo burbuja tardo: ", final)
	canal1 <- 1
	//fmt.Println(arreglo)
	return arreglo
}

func ordenamientoSeleccion(canal1 chan int, arreglo []int) []int {
	inicio := time.Now()
	//fmt.Println("El metodo inicio: ", inicio)
	for i := 0; i < len(arreglo)-1; i++ {
		min := i
		for j := i + 1; j < len(arreglo); j++ {
			if arreglo[j] < arreglo[min] {
				arreglo[j], arreglo[min] = arreglo[min], arreglo[j]
			}
		}
	}
	final := time.Since(inicio)
	fmt.Println("El metodo seleccion tardo: ", final)
	canal1 <- 1
	//fmt.Println(arreglo)
	return arreglo
}
