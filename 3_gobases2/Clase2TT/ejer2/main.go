package main

import "fmt"

type Matrix struct {
	Columns     int
	Rows        int
	Maximum     float64
	IsCuadratic bool
	Data        [][]float64
}

func (m *Matrix) Set(matrix [][]float64) {

	m.Data = matrix
	m.IsCuadratic = true
	m.Columns = len(matrix)

	for _, row := range matrix {

		m.Rows = len(row)

		if len(row) != 3 && m.Rows != 1 {
			m.IsCuadratic = false
		}

		for _, value := range row {

			if m.Maximum < value {
				m.Maximum = value
			}
		}
	}
}

func (m Matrix) Detail() {

	for _, row := range m.Data {

		for _, value := range row {
			fmt.Printf("%.2f ", value)
		}

		fmt.Print("\n")
	}
	fmt.Printf("Columns: %d\nRows: %d\nMaximum Value: %.2f\nIs Cuadratic: %v", m.Columns, m.Rows, m.Maximum, m.IsCuadratic)
}

func main() {

	matrix := Matrix{}

	var matrixData = [][]float64{
		{0.1, 0.2, 0.3},
		{0.4, 0.5, 0.6},
		{0.7, 0.8, 0.9},
	}

	matrix.Set(matrixData)
	matrix.Detail()
}
