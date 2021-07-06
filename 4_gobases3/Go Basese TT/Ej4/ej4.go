package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	primerLista := rand.Perm(1000)
	//segundaLista := rand.Perm(1000)
	//tercerLista := rand.Perm(1000000)
	canalIncercion, canalBurbuja, canalSeleccion := make(chan time.Time), make(chan time.Time), make(chan time.Time)
	copia1, copia2, copia3 := multiplicarLista(primerLista)

	inicio := time.Now()
	go OrdenamientoIncercion(copia1, canalIncercion)
	go OrdenamientoBurbuja(copia2, canalBurbuja)
	go OrdenamientoSeleccion(copia3, canalSeleccion)

	finIncercion := <-canalIncercion
	finBurbuja := <-canalBurbuja
	finSeleccion := <-canalSeleccion

	fmt.Println("Seleccion: ", finSeleccion.Sub(inicio))
	fmt.Println("Burbuja: ", finBurbuja.Sub(inicio))
	fmt.Println("Incercion: ", finIncercion.Sub(inicio))

	//go OrdenamientoSeleccion(primerLista, canalSeleccion)

}

func multiplicarLista(lista []int) ([]int, []int, []int) {
	copia1, copia2, copia3 := make([]int, len(lista)), make([]int, len(lista)), make([]int, len(lista))
	copy(copia1, lista)
	copy(copia2, lista)
	copy(copia3, lista)
	return copia1, copia2, copia3
}
func OrdenamientoIncercion(lista []int, c chan time.Time) {
	var n = len(lista)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if lista[j-1] > lista[j] {
				lista[j-1], lista[j] = lista[j], lista[j-1]
			}
			j = j - 1
		}
	}
	c <- time.Now()
}
func OrdenamientoBurbuja(lista []int, c chan time.Time) {
	for i := len(lista); i > 0; i-- {
		for j := 1; j < i; j++ {
			if lista[j-1] > lista[j] {
				aux := lista[j]
				lista[j] = lista[j-1]
				lista[j-1] = aux
			}
		}
	}
	c <- time.Now()
}
func OrdenamientoSeleccion(lista []int, c chan time.Time) {
	var n = len(lista)
	for i := 0; i < n; i++ {
		var indice = i
		for j := i; j < n; j++ {
			if lista[j] < lista[indice] {
				indice = j
			}
		}
		lista[i], lista[indice] = lista[indice], lista[i]
	}
	c <- time.Now()
}
