// Ejercicio 2 - Matrix
// Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que represente una matriz de datos.
// Para ello requieren una estructura Matrix que tenga los métodos:
// Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
// Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
// La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho, si es cuadrática y cuál es el valor máximo.

package main

import (
	f "fmt"
)

type Matrix struct {
	height       int
	width        int
	isCuadratic  bool
	maxValue     float64
	matrizValues [][]float64
}

func (m *Matrix) set(h int, w int, matrizValues []float64) {
	m.height = h
	m.width = w

	if m.height != m.width {
		m.isCuadratic = true
	} else {
		m.isCuadratic = false
	}

	m.maxValue = matrizValues[0]
	posAux := 0
	for i := 0; i < m.height; i++ {
		m.matrizValues[i] = make([]float64, w)
		for j := 0; j < m.width; j++ {
			m.matrizValues[i][j] = matrizValues[posAux]
			if matrizValues[posAux] > m.maxValue {
				m.maxValue = matrizValues[posAux]
			}
			posAux++
		}
	}
}

func (m *Matrix) print() {

	for i := 0; i < m.height; i++ {
		for j := 0; j < m.width; j++ {
			f.Printf("|%10.2f|", m.matrizValues[i][j])
		}
		f.Println()
	}
}

func main() {

	var matrizValues []float64
	var height int
	var width int
	var value float64

	f.Printf("Ingrese el alto de la matriz: ")
	f.Scan(&height)
	f.Printf("Ingrese el ancho de la matriz: ")
	f.Scan(&width)

	for width <= 0 || height <= 0 {
		f.Printf("Reingrese el alto de la matriz: ")
		f.Scan(&height)
		f.Printf("Reingrese el ancho de la matriz: ")
		f.Scan(&width)
	}

	var matriz1 = Matrix{
		matrizValues: make([][]float64, height),
	}

	for i := 0; i < width*height; i++ {
		f.Printf("Ingrese el valor %d de la matriz: ", i+1)
		f.Scan(&value)
		matrizValues = append(matrizValues, value)
	}

	matriz1.set(height, width, matrizValues)
	matriz1.print()
	f.Printf("El valor más grande dentro de la matriz es: %.2f\n", matriz1.maxValue)
}
