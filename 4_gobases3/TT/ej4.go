package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Burbuja(ListaDesordenada []int, cb chan float64) {
	start := time.Now()

	var auxiliar int
	for i := 0; i < len(ListaDesordenada); i++ {
		for j := 0; j < len(ListaDesordenada); j++ {
			if ListaDesordenada[i] < ListaDesordenada[j] {
				auxiliar = ListaDesordenada[i]
				ListaDesordenada[i] = ListaDesordenada[j]
				ListaDesordenada[j] = auxiliar
			}
		}
	}
	elapsed := time.Since(start)
	cb <- float64(elapsed.Milliseconds())
}

func Insercion(ListaDesordenada []int, ci chan float64) {
	start := time.Now()

	var auxiliar int
	for i := 1; i < len(ListaDesordenada); i++ {
		auxiliar = ListaDesordenada[i]
		for j := i - 1; j >= 0 && ListaDesordenada[j] > auxiliar; j-- {
			ListaDesordenada[j+1] = ListaDesordenada[j]
			ListaDesordenada[j] = auxiliar
		}
	}
	elapsed := time.Since(start)
	ci <- float64(elapsed.Milliseconds())
}

func Seleccion(ListaDesordenada []int, cs chan float64) {
	start := time.Now()

	var n = len(ListaDesordenada)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if ListaDesordenada[j] < ListaDesordenada[minIdx] {
				minIdx = j
			}
		}
		ListaDesordenada[i], ListaDesordenada[minIdx] = ListaDesordenada[minIdx], ListaDesordenada[i]
	}
	elapsed := time.Since(start)
	cs <- float64(elapsed.Milliseconds())
}

func main() {

	variable1 := rand.Perm(100)

	cs := make(chan float64)
	ci := make(chan float64)
	cb := make(chan float64)

	go Seleccion(variable1, cs)
	go Burbuja(variable1, cb)
	go Insercion(variable1, ci)

	var resS = -1.0
	var resB = -1.0
	var resI = -1.0
	resS = <-cs
	resB = <-cb
	resI = <-ci

	fmt.Println("100 elementos con Seleccion", resS)
	fmt.Println("100 elementos con Burbuja", resB)
	fmt.Println("100 elementos con Insercion", resI)

	variable2 := rand.Perm(10000)

	if resS != -1.0 && resB != -1.0 && resI != -1.0 {
		go Seleccion(variable2, cs)
		go Burbuja(variable2, cb)
		go Insercion(variable2, ci)

		var resS2 = -1.0
		var resB2 = -1.0
		var resI2 = -1.0
		resS2 = <-cs
		resB2 = <-cb
		resI2 = <-ci

		fmt.Println("10000 elementos con Seleccion", resS2)
		fmt.Println("10000 elementos con Burbuja", resB2)
		fmt.Println("10000 elementos con Insercion", resI2)

		if resS2 != -1.0 && resB2 != -1.0 && resI2 != -1.0 {
			variable3 := rand.Perm(1000000)

			go Seleccion(variable3, cs)
			go Burbuja(variable3, cb)
			go Insercion(variable3, ci)

			resS3 := <-cs
			resB3 := <-cb
			resI3 := <-ci

			fmt.Println("1000000 elementos con Seleccion", resS3)
			fmt.Println("1000000 elementos con Burbuja", resB3)
			fmt.Println("1000000 elementos con Insercion", resI3)

		}
	}
}
