package main

import "fmt"

func main() {
	/*
		Ejercicio 2 - Matrix
		Una empresa de inteligencia artificial necesita tener una funcionalidad para crear matrices en
		Go.
		Para ello requieren una estructura Matrix que tenga los métodos:
		- Set: Recibe una matrix de flotantes e inicializa los valores en la estructura Matrix
		- Print: Imprime por pantalla la matrix de una formas más visible (Con los saltos de línea
		entre filas)
		La estructura matrix debe guardar la matrix, el tamaño de alto, el tamaño de ancho, si es
		cuadrática y el valor máximo.
	*/

	matrix := Matrix{}
	matrix.Set([][]int64{{1, 1, 1}, {4, 1, 7}, {1, 65, 1}})
	matrix.Print()
	fmt.Println("ALTO:", matrix.Alto, "\nANCHO:", matrix.Ancho, "\nCUADRATICA:", matrix.Cuadratica, "\nMAX:", matrix.Max)
}

type Matrix struct {
	Matrix     [][]int64
	Alto       int64
	Ancho      int64
	Cuadratica bool
	Max        int64
}

func (m *Matrix) Set(matrix [][]int64) {

	for i, submatrix := range matrix {
		for j, value := range submatrix {
			if i == 0 && j == 0 {
				m.Max = value
			} else {
				if m.Max < value {
					m.Max = value
				}
			}
			m.Ancho = int64(j + 1)
		}
		m.Alto = int64(i + 1)
	}
	m.Matrix = matrix
	m.Cuadratica = false
}

func (m *Matrix) Print() {

	for _, row := range m.Matrix {
		for j, value := range row {
			if j != int(m.Ancho-1) {
				fmt.Printf("%d ", value)
			} else {
				fmt.Printf("%d\n", value)
			}
		}
	}
}
