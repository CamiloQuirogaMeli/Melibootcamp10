package main

import "fmt"

type Matrix struct {
	Matrix      [][]float64
	Alto        int
	Ancho       int
	ValorMaximo float64
	Cuadratica  bool
}

func (m *Matrix) Set(matrix [][]float64) {
	m.Matrix = matrix
	m.Alto = obtenerAlto(matrix)
	m.Ancho = len(matrix)
	m.ValorMaximo = maxValue(matrix)
	m.Cuadratica = cuadratica(m.Alto, m.Ancho)
}

func obtenerAlto(matrix [][]float64) int {
	alto := len(matrix[0])
	for _, element := range matrix {
		if len(element) > alto {
			alto = len(element)
		}
	}
	return alto
}

func maxValue(matrix [][]float64) float64 {
	var maxValue float64
	for _, element := range matrix {
		for _, element := range element {
			if maxValue < element {
				maxValue = element
			}
		}
	}
	return maxValue
}

func (m Matrix) Print() {
	for _, element := range m.Matrix {
		fmt.Println(element)
	}
}

func cuadratica(alto int, ancho int) bool {
	if (alto == ancho) && alto != 0 {
		return true
	}
	return false
}

func main() {
	m := Matrix{}
	m.Set([][]float64{{0, 0, 4}, {1, 2}, {2, 4}, {3, 6, 5, 8}, {4, 8, 5, 6, 7, 17}})
	m.Print()
}
