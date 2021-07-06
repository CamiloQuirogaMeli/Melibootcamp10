package main

import "fmt"

type Matrix struct {
	matrix   [][]float64
	height   int
	width    int
	square   bool
	maxValue float64
}

func (m *Matrix) Set(mat [][]float64) {
	m.matrix = mat

	m.width = len(m.matrix[0])
	m.height = len(m.matrix)

	m.square = m.height == m.width

	var max = m.matrix[0][0]

	for _, row := range m.matrix {
		for _, value := range row {
			if max < value {
				max = value
			}
		}
	}

	m.maxValue = max
}

func (m Matrix) Print() {
	for _, row := range m.matrix {
		fmt.Println(row)
	}
}

func main() {

	row1 := []float64{1.1, 1.2, 1.3, 1.4}
	row2 := []float64{2.1, 2.2, 2.3, 2.4}
	row3 := []float64{3.1, 3.2, 3.3, 3.4}
	row4 := []float64{4.1, 4.2, 4.3, 4.4}

	mat := [][]float64{row1, row2, row3, row4}

	matrix := Matrix{}

	matrix.Set(mat)

	matrix.Print()
}
