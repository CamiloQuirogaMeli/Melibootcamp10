package ejercicio2

import "fmt"

type Matrix struct {
	Matrix      [][]float64
	Height      uint
	Width       uint
	MaxValue    float64
	IsQuadratic bool
}

func (m *Matrix) SetMatrix(matrix [][]float64) {
	m.Matrix = matrix
	m.Height = uint(len(matrix))
	m.Width = uint(len(matrix[0]))
	m.IsQuadratic = m.Height == m.Width
	m.calculateMaxValue()
}

func InitializeFloatMatrix(height uint, width uint) [][]float64 {
	m := make([][]float64, height)
	for i := uint(0); i < height; i++ {
		m[i] = make([]float64, width)
		for j := uint(0); j < width; j++ {
			m[i][j] = float64(i*j + 1)
		}
	}
	return m
}

func (m Matrix) PrintMatrix() {
	for i := uint(0); i < m.Height; i++ {
		fmt.Println(m.Matrix[i])
	}
}

func (m *Matrix) calculateMaxValue() {
	for i := uint(0); i < m.Height; i++ {
		for j := uint(0); j < m.Width; j++ {
			if m.MaxValue < m.Matrix[i][j] {
				m.MaxValue = m.Matrix[i][j]
			}
		}
	}
}
