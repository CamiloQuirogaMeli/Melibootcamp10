// Ejercicio 4 - Ordenamiento
// Una empresa de sistemas requiere analizar qué algoritmos de ordenamiento utilizar para sus servicios.
// Para ellos se requiere instanciar 3 arreglos con valores aleatorios desordenados
// un arreglo de números enteros con 100 valores
// un arreglo de números enteros con 1.000 valores
// un arreglo de números enteros con 1.000.000 valores

// Para instanciar las variables utilizar rand
// package main

// import (
//    "math/rand"
// )

// func main() {
//    variable1 := rand.Perm(100)
//    variable2 := rand.Perm(1000)
//    variable3 := rand.Perm(1000000)
// }

// Se debe realizar el ordenamiento de cada una por:
// Ordenamiento por inserción
// Ordenamiento por burbuja
// Ordenamiento por selección

// Una go routine por cada ejecución de ordenamiento
// Debo esperar a que terminen los ordenamientos de 100 números para seguir el de 1.000 y después el de 1.000.000.
// Por último debo medir el tiempo de cada uno y mostrar en pantalla el resultado, para saber qué ordenamiento fue mejor para cada arreglo

package main

import (
	f "fmt"
	"math/rand"
	"time"
)

type chanInfoOrd struct {
	nameOrd string
	timeOrd float64
}

func ordInser(numbers []int, c chan chanInfoOrd) {

	oldTime := time.Now()
	var aux int
	for i := 1; i < len(numbers); i++ {
		aux = numbers[i]
		for j := i - 1; j >= 0 && numbers[j] > aux; j-- {
			numbers[j+1] = numbers[j]
			numbers[j] = aux
		}
	}

	currentTime := time.Now()
	diff := currentTime.Sub(oldTime)

	var valueChan chanInfoOrd = chanInfoOrd{"Insertion", diff.Seconds()}

	c <- valueChan
}

func ordBubble(numbers []int, c chan chanInfoOrd) {

	oldTime := time.Now()
	var aux int
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if numbers[i] > numbers[j] {
				aux = numbers[i]
				numbers[i] = numbers[j]
				numbers[j] = aux
			}
		}
	}
	currentTime := time.Now()
	diff := currentTime.Sub(oldTime)

	var valueChan chanInfoOrd = chanInfoOrd{"Bubble", diff.Seconds()}

	c <- valueChan
}

func ordSelect(numbers []int, c chan chanInfoOrd) {

	oldTime := time.Now()

	var aux int
	for i := 0; i < len(numbers); i++ {
		aux = i
		for j := i + 1; j < len(numbers); j++ {
			if numbers[aux] > numbers[j] {
				aux = j
			}
		}
		temp := numbers[i]
		numbers[i] = numbers[aux]
		numbers[aux] = temp
	}

	currentTime := time.Now()
	diff := currentTime.Sub(oldTime)

	var valueChan chanInfoOrd = chanInfoOrd{"Selection", diff.Seconds()}

	c <- valueChan

}

func main() {

	c := make(chan chanInfoOrd)
	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(1000000)

	go ordBubble(variable1, c)
	go ordSelect(variable1, c)
	go ordInser(variable1, c)

	result := <-c
	f.Printf("Finalizó el ordenamiento %s con %d valores en la lista y duró (en segundos): %f\n", result.nameOrd, len(variable1), result.timeOrd)
	result = <-c
	f.Printf("Finalizó el ordenamiento %s con %d valores en la lista y duró (en segundos): %f\n", result.nameOrd, len(variable1), result.timeOrd)
	result = <-c
	f.Printf("Finalizó el ordenamiento %s con %d valores en la lista y duró (en segundos): %f\n", result.nameOrd, len(variable1), result.timeOrd)

	f.Println()

	go ordBubble(variable2, c)
	go ordSelect(variable2, c)
	go ordInser(variable2, c)

	result = <-c
	f.Printf("Finalizó el ordenamiento %s con %d valores en la lista y duró (en segundos): %f\n", result.nameOrd, len(variable2), result.timeOrd)
	result = <-c
	f.Printf("Finalizó el ordenamiento %s con %d valores en la lista y duró (en segundos): %f\n", result.nameOrd, len(variable2), result.timeOrd)
	result = <-c
	f.Printf("Finalizó el ordenamiento %s con %d valores en la lista y duró (en segundos): %f\n", result.nameOrd, len(variable2), result.timeOrd)

	f.Println()

	go ordBubble(variable3, c)
	go ordSelect(variable3, c)
	go ordInser(variable3, c)

	result = <-c
	f.Printf("Finalizó el ordenamiento %s con %d valores en la lista y duró (en segundos): %.2f\n", result.nameOrd, len(variable3), result.timeOrd)
	result = <-c
	f.Printf("Finalizó el ordenamiento %s con %d valores en la lista y duró (en segundos): %.2f\n", result.nameOrd, len(variable3), result.timeOrd)
	result = <-c
	f.Printf("Finalizó el ordenamiento %s con %d valores en la lista y duró (en segundos): %.2f\n", result.nameOrd, len(variable3), result.timeOrd)

}
