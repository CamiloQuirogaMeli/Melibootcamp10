package main

import (
	"fmt"
)

const (
	alto  = 3
	ancho = 5
)

type Matrix struct {
	alto         int
	ancho        int
	matrix       [alto][ancho]float64
	esCuadratica bool
	maximo       float64
}

func (m *Matrix) Set(numInicial float64) {
	var matrix [alto][ancho]float64
	var esCuadratica bool
	for i := 0; i < alto; i++ {
		for j := 0; j < ancho; j++ {
			matrix[i][j] = numInicial
		}
	}

	if ancho == alto {
		esCuadratica = true
	} else {
		esCuadratica = false
	}

	m.matrix = matrix
	m.esCuadratica = esCuadratica
	m.maximo = numInicial
	m.alto = alto
	m.ancho = ancho
	fmt.Println("Matrix inicializada")
}

func (m Matrix) Print() {
	fmt.Println("Alto de la matriz:", m.alto)
	fmt.Println("Ancho de la matriz:", m.ancho)
	fmt.Println("Número máximo:", m.maximo)

	fmt.Println()
	fmt.Print("-")
	for l := 0; l < m.ancho; l++ {
		fmt.Print("-")
	}

	fmt.Println("-")
	for i := 0; i < m.alto; i++ {
		fmt.Print("|")
		for j := 0; j < m.ancho; j++ {
			fmt.Print(m.matrix[i][j])
		}
		fmt.Println("|")
	}
	fmt.Print("-")
	for l := 0; l < m.ancho; l++ {
		fmt.Print("-")
	}
	fmt.Println("-")
}

func main() {
	/*
			Una empresa de inteligencia artificial necesita tener una funcionalidad para crear matrices en
		Go.
		Para ello requieren una estructura Matrix que tenga los métodos:
		- Set: Recibe una matrix de flotantes e inicializa los valores en la estructura Matrix
		- Print: Imprime por pantalla la matrix de una formas más visible (Con los saltos de línea
		entre filas)
		La estructura matrix debe guardar la matrix, el tamaño de alto, el tamaño de ancho, si es
		cuadrática y el valor máximo.
	*/

	var accion float64
	salir := false
	numInicial := 0.0

	matrix := Matrix{}

	for !salir {
		fmt.Println("Digita la accion a realizar:")
		fmt.Println("	1: Ejecutar el método set")
		fmt.Println("	2: Ejecutar el método print")
		fmt.Println("	3: Salir")
		fmt.Scanln(&accion)
		switch accion {
		case 1:
			fmt.Println("Digita el numero inicial:")
			fmt.Scanln(&numInicial)
			matrix.Set(numInicial)
		case 2:
			matrix.Print()
		case 3:
			salir = true
		default:
			fmt.Println("Opción no valida")
		}

	}
}
