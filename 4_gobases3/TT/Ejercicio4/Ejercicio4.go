package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
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

	var1 := rand.Perm(100)
	var2 := rand.Perm(1000)
	var3 := rand.Perm(1000000)

	chanel1 := make(chan time.Duration)
	chanel2 := make(chan time.Duration)
	chanel3 := make(chan time.Duration)

	go sortInsert(var1, chanel1)
	go sortBubble(var1, chanel2)
	go sortSelect(var1, chanel3)

	fmt.Println("Variable 1 - 100 elementos")
	fmt.Printf("Insercion %d ", <-chanel1)
	fmt.Printf("Burbuja %d ", <-chanel2)
	fmt.Printf("Seleccion %d ", <-chanel3)

	go sortInsert(var2, chanel1)
	go sortBubble(var2, chanel2)
	go sortSelect(var2, chanel3)

	fmt.Println("Variable 1 - 1000 elementos")
	fmt.Printf("Insercion %d ", <-chanel1)
	fmt.Printf("Burbuja %d ", <-chanel2)
	fmt.Printf("Seleccion %d ", <-chanel3)

	go sortInsert(var3, chanel1)
	go sortBubble(var3, chanel2)
	go sortSelect(var3, chanel3)

	fmt.Println("Variable 1 - 1000000 elementos")
	fmt.Printf("Insercion %d ", <-chanel1)
	fmt.Printf("Burbuja %d ", <-chanel2)
	fmt.Printf("Seleccion %d ", <-chanel3)

}

func sortInsert(a []int, c chan time.Duration) {
	s := time.Now()
	var i int
	var j int
	var aux int

	for i = 1; i < len(a); i++ {
		aux = a[i]
		j = i - 1
		for j >= 0 && aux < a[j] {
			a[j+1] = a[j]
		}
		a[j+1] = aux
	}

	e := time.Now()
	t := e.Sub(s)
	c <- t
}

func sortBubble(a []int, c chan time.Duration) {
	s := time.Now()
	var i int
	var j int
	var aux int

	for i = 0; i < len(a)-1; j++ {

		for j = 0; j < len(a)-i-1; j++ {
			if a[j+1] < a[j] {
				aux = a[j+1]
				a[j+1] = a[j]
				a[j] = aux
			}
		}
	}
	e := time.Now()
	t := e.Sub(s)
	c <- t
}

func sortSelect(a []int, c chan time.Duration) {
	s := time.Now()
	var i int
	var j int
	var min int
	var aux int

	for i = 0; i < len(a)-1; i++ {
		min = i
		for j = i + 1; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		if i != min {
			aux = a[i]
			a[i] = a[min]
			a[min] = aux
		}
	}

	e := time.Now()
	t := e.Sub(s)
	c <- t
}
