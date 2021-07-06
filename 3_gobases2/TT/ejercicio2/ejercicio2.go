package main

import (
	"errors"
	"fmt"
	"math"
)

type Matrix struct {
	matrix [][]float64
	height int
	width  int
	square bool
	maxVal float64
}

func (m *Matrix) setMatrix(newMatrix [][]float64) error {
	m.height = len(newMatrix)
	if m.height > 0 && len(newMatrix[0]) > 0 {
		m.width = len(newMatrix[0])
	} else {
		errors.New("Debe tener al menos un valor")
	}
	if m.height == m.width {
		m.square = true
	} else {
		m.square = false
	}
	m.matrix = newMatrix
	return nil
}

func (m *Matrix) findMax() {
	maxVal := -math.MaxFloat64

	for i := 0; i < m.height; i++ {
		for j := 0; j < m.width; j++ {
			maxVal = math.Max(maxVal, m.matrix[i][j])
		}
	}
	m.maxVal = maxVal
}

func (m Matrix) printMatrix() {
	for i := 0; i < m.height; i++ {
		row := ""
		for j := 0; j < m.width; j++ {
			row += fmt.Sprintf("%.2f\t", m.matrix[i][j])
		}
		fmt.Println(row)
	}
}

func main() {
	m := Matrix{}
	matrix := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	m.setMatrix(matrix)
	m.findMax()
	m.printMatrix()
	fmt.Printf("Numero Filas: %d\n", m.height)
	fmt.Printf("Numero Columnas: %d\n", m.width)
	fmt.Printf("Es cuadrada: %t\n", m.square)
	fmt.Printf("El valor mas grande en la matriz es %.2f\n", m.maxVal)
}
