package main

import "fmt"

type Matrix struct {
	matrix      [][]float64
	height      int
	width       int
	isCuadratic bool
	maxValue    float64
}

func (m *Matrix) Set(value float64) {
	m.width = len(m.matrix)
	m.height = len(m.matrix[0])
	m.isCuadratic = m.width == m.height
	m.maxValue = 0.0
	for _, h := range m.matrix {

		for i := range h {
			if h[i] > m.maxValue {
				m.maxValue = h[i]
			}
			h[i] = value
		}
	}
}
func (m Matrix) Print() {
	for _, h := range m.matrix {
		for i := range h {
			fmt.Printf("[%.2f] ", h[i])
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("Ejercicio 2")
	myMatrix := Matrix{
		matrix: [][]float64{
			{0, 1, 2, 3},
			{4, 5, 6, 7},
			{8, 9, 9.9, 7.0},
			{1, 3, 6, 7},
		},
	}
	fmt.Println("Matriz original:")
	myMatrix.Print()
	myMatrix.Set(1.00)
	fmt.Println("¿Es cuadratica la matriz? -", myMatrix.isCuadratic)
	fmt.Println("El maximo valor de la matriz es:", myMatrix.maxValue)
	fmt.Println("\nMatriz luego de setear todos los campos en 1.00")
	myMatrix.Print()
}
