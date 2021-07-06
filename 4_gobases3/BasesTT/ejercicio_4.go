package main

import (
	"fmt"
	"math/rand"
	"time"
)

func insercion(datos []int) {
	for i := 0; i < len(datos); i++ {
		clave := datos[i]
		j := i - 1
		for j >= 0 && datos[j] > clave {
			datos[j+1] = datos[j]
			j--
		}
		datos[j+1] = clave
	}
}

func burbuja(datos []int) {
	for i := 0; i < len(datos); i++ {
		for j := 0; j < len(datos)-1; j++ {
			if datos[j] > datos[j+1] {
				auxiliar := datos[j+1]
				datos[j+1] = datos[j]
				datos[j] = auxiliar
			}
		}
	}
}

func seleccion(datos []int) {
	for i := 0; i < len(datos); i++ {
		minimo := i
		for j := i + 1; j < len(datos); j++ {
			if datos[j] < datos[minimo] {
				minimo = j
			}
		}
		if i != minimo {
			auxiliar := datos[i]
			datos[i] = datos[minimo]
			datos[minimo] = auxiliar
		}
	}
}

func medirTiempo(funcion func([]int), datos []int, id string, finalizado chan string) {
	inicio := time.Now()
	funcion(datos)
	duracion := time.Since(inicio)
	finalizado <- fmt.Sprintf("Tiempo de ejecución %v (%d): %v\n", id, len(datos), duracion)
}

func main() {
	respuesta := make(chan string)
	datos := [][]int{rand.Perm(100), rand.Perm(1000), rand.Perm(10000)}

	for i := 0; i < len(datos); i++ {
		go medirTiempo(insercion, datos[i], "insercion", respuesta)
		go medirTiempo(burbuja, datos[i], "burbuja", respuesta)
		go medirTiempo(seleccion, datos[i], "seleccion", respuesta)
		for j := 0; j < 3; j++ {
			fmt.Print(<-respuesta)
		}
	}
}
