package main

import "fmt"

/*
	Ejercicio 2 - Matrix
	Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una
	estructura que represente una matriz de datos.
	Para ello requieren una estructura Matrix que tenga los métodos:
	- Set: Recibe una serie de valores de punto flotante e inicializa los valores en la
	estructura Matrix
	- Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea
	entre filas)
	La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la
	dimensión del ancho, si es cuadrática y cuál es el valor máximo.
*/

type Matrix struct {
	data       [][]float64
	alto       int
	ancho      int
	cuadratica bool
	valorMax   float64
}

func (m *Matrix) Set(matrix [][]float64) {
	m.data = matrix
	m.alto = len(matrix)
	m.ancho = len(matrix[0])
	m.cuadratica = (m.alto == m.ancho)
	m.valorMax = 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if m.valorMax < matrix[i][j] {
				m.valorMax = matrix[i][j]
			}
		}
	}
}

func (m Matrix) Print() {
	for i := 0; i < len(m.data); i++ {
		for j := 0; j < len(m.data[i]); j++ {
			fmt.Print(m.data[i][j], "  ")
		}
		fmt.Print("\n")
	}
	fmt.Println("Alto: ", m.alto)
	fmt.Println("Ancho: ", m.ancho)
	if m.cuadratica {
		fmt.Println("Es cuadratica")
	} else {
		fmt.Println("No es cuadratica")
	}
	fmt.Println("Valor maximo: ", m.valorMax)
}

func main() {
	var m Matrix
	columnas := []float64{10.1, 10.2, 10.3}
	filas := []float64{11.1, 11.2, 11.3}
	matrix := [][]float64{columnas, filas}
	m.Set(matrix)
	m.Print()
}
