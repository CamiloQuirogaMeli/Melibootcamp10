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

variable1 := rand.Perm(100)
variable2 := rand.Perm(1000)
variable3 := rand.Perm(1000000)

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

func porInserción(datos []int, c chan []int) {
	tInicio := time.Now()
	aux := 0
	for i := 1; i < len(datos); i++ {
		j := i
		for j > 0 && datos[j] < datos[j-1] {
			aux = datos[j-1]
			datos[j-1] = datos[j]
			datos[j] = aux
			j--
		}
	}
	tFinal := time.Since(tInicio)
	fmt.Println("Tiempo inserción: ", tFinal)
	c <- datos
}

func porBurbuja(datos []int, c chan []int) {
	tInicio := time.Now()
	aux := 0
	for i := 0; i < len(datos)-1; i++ {
		for j := 0; j < len(datos)-i-1; j++ {
			if datos[j] > datos[j+1] {
				aux = datos[j]
				datos[j] = datos[j+1]
				datos[j+1] = aux
			}
		}
	}
	tFinal := time.Since(tInicio)
	fmt.Println("Tiempo burbuja: ", tFinal)
	c <- datos
}

func porSeleccion(datos []int, c chan []int) {
	tInicio := time.Now()
	var min, pos int
	for i := 0; i < len(datos)-1; i++ {
		min = datos[i]
		pos = i
		for j := i + 1; j < len(datos); j++ {
			if datos[j] < min {
				min = datos[j]
				pos = j
			}
		}
		datos[pos] = datos[i]
		datos[i] = min
	}
	tFinal := time.Since(tInicio)
	fmt.Println("Tiempo selección: ", tFinal)
	c <- datos

}

func main() {

	c1 := make(chan []int)
	c2 := make(chan []int)
	c3 := make(chan []int)

	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(1000000)

	/*ARREGLO DE 100 ELEMENTOS*/
	fmt.Println("ARREGLO DE 100 ELEMENTOS")
	go porInserción(variable1, c1)
	go porBurbuja(variable1, c2)
	go porSeleccion(variable1, c3)
	<-c1
	<-c2
	<-c3

	/*ARREGLO DE 1000 ELEMENTOS*/
	fmt.Println("ARREGLO DE 1000 ELEMENTOS")
	go porInserción(variable2, c1)
	go porBurbuja(variable2, c2)
	go porSeleccion(variable2, c3)
	<-c1
	<-c2
	<-c3

	/*ARREGLO DE 1000000 DE ELEMENTOS*/
	fmt.Println("ARREGLO DE 1000000 ELEMENTOS")
	go porInserción(variable3, c1)
	go porBurbuja(variable3, c2)
	go porSeleccion(variable3, c3)
	<-c1
	<-c2
	<-c3
}
