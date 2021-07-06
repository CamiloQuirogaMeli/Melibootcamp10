package main

import (
	"fmt"
	"math/rand"
	"time"
)

func InsercionSort(num []int) {
	var n = len(num)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if num[j-1] > num[j] {
				num[j-1], num[j] = num[j], num[j-1]
			}
			j = j - 1
		}
	}
}
func BubbleSort(num []int) {
	n := len(num)
	cambio := true
	for cambio {
		cambio = false
		for i := 0; i < n-1; i++ {

			if num[i] > num[i+1] {

				num[i], num[i+1] = num[i+1], num[i]

				cambio = true
			}
		}
	}
}

func SelectionSort(num []int) {
	var n = len(num)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if num[j] < num[minIdx] {
				minIdx = j
			}
		}
		num[i], num[minIdx] = num[minIdx], num[i]
	}
}

func MedirTiempo(funcion func(num []int), num []int, c chan time.Duration) {
	i := time.Now()
	funcion(num)
	//fmt.Println("Tiempo de ejecución:", time.Since(i))
	c <- time.Since(i)
}

func main() {
	/*

		Una empresa de sistemas requiere analizar qué algoritmos de ordenamiento utilizar para sus
		servicios.
		Para ellos se requiere instanciar 3 arreglos con valores aleatorios desordenados
				- un arreglo de números enteros con 100 valores
				- un arreglo de números enteros con 1.000 valores
				- un arreglo de números enteros con 1.000.000 valores
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
	numeros := make([][]int, 0)

	numeros = append(numeros, rand.Perm(100))
	numeros = append(numeros, rand.Perm(1000))
	numeros = append(numeros, rand.Perm(100000))

	fmt.Println("Números creados")
	c := make(chan time.Duration) //Indican el fin de la ejecución

	for i := 0; i < 3; i++ {
		fmt.Println("\nTiempos para", len(numeros[i]), "\n")

		go MedirTiempo(BubbleSort, numeros[i], c)
		fmt.Println("Tiempo BubbleSort ", <-c)

		go MedirTiempo(InsercionSort, numeros[i], c)
		fmt.Println("Tiempo InsercionSort ", <-c)

		go MedirTiempo(SelectionSort, numeros[i], c)
		fmt.Println("Tiempo SelectionSort ", <-c)

	}

}
