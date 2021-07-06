package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan []int)
	variableCien := rand.Perm(100)
	variableMil := rand.Perm(1000)
	variableMillon := rand.Perm(1000000)
	fmt.Println("*** ARREGLO DE CIEN ***")
	go ordInsercion(variableCien, c)
	go ordBurbuja(variableCien, c)
	go ordSeleccion(variableCien, c)
	for i := 0; i < 3; i++ {
		<-c
	}

	fmt.Println("*** ARREGLO DE MIL ***")
	go ordInsercion(variableMil, c)
	go ordBurbuja(variableMil, c)
	go ordSeleccion(variableMil, c)
	for i := 0; i < 3; i++ {
		<-c
	}

	fmt.Println("*** ARREGLO DE UN MILLÓN ***")
	go ordInsercion(variableMillon, c)
	go ordBurbuja(variableMillon, c)
	go ordSeleccion(variableMillon, c)
	for i := 0; i < 3; i++ {
		<-c
	}
}

func ordInsercion(arr2 []int, c chan []int) {
	arr := arr2
	var aux int
	comienzo := time.Now()
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			aux = arr[j]
			arr[j] = arr[j-1]
			arr[j-1] = aux
			j--
		}
	}
	tiempoTranscurrido := time.Since(comienzo)
	fmt.Println("Tiempo transcurrido en ordenamiento por inserción: ", tiempoTranscurrido)
	c <- arr
}

func ordBurbuja(arr2 []int, c chan []int) {
	arr := arr2
	var aux int
	comienzo := time.Now()
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j+1] < arr[j] {
				aux = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = aux
			}
		}
	}
	tiempoTranscurrido := time.Since(comienzo)
	fmt.Println("Tiempo transcurrido en ordenamiento por burbuja: ", tiempoTranscurrido)
	c <- arr
}

func ordSeleccion(arr2 []int, c chan []int) {
	arr := arr2
	var aux int
	comienzo := time.Now()
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				aux = arr[j]
				arr[j] = arr[i]
				arr[i] = aux
			}
		}
	}
	tiempoTranscurrido := time.Since(comienzo)
	fmt.Println("Tiempo transcurrido en ordenamiento por selección: ", tiempoTranscurrido)
	c <- arr
}
