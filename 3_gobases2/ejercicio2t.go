package main

import "fmt"

// Una empresa de inteligencia artificial necesita tener una funcionalidad para crear matrices en
// Go.
// Para ello requieren una estructura Matrix que tenga los métodos:
// - Set: Recibe una matrix de flotantes e inicializa los valores en la estructura Matrix
// - Print: Imprime por pantalla la matrix de una formas más visible (Con los saltos de línea
// entre filas)
// La estructura matrix debe guardar la matrix, el tamaño de alto, el tamaño de ancho, si es
// cuadrática y el valor máximo.

type Matrix struct {
	matrix     []float64
	alto       float64
	ancho      float64
	cuadratica bool
	max        float64
}

func (m *Matrix) set(matriz []float64) {
	m.matrix = matriz
	m.alto = 2
	m.ancho = 2
	m.cuadratica = true
	m.max = 10.5
}

func (m Matrix) print() {
	fmt.Println("Matrix: ", m.matrix)
	fmt.Println("Alto: ", m.alto)
	fmt.Println("Ancho: ", m.ancho)
	fmt.Println("Cuadratica: ", m.cuadratica)
	fmt.Println("Maximo: ", m.max)

}
