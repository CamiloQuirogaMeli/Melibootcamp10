package main

import "fmt"

func main() {
	test := Matrix{}
	test.Set([][]int64{{1, 1, 1}, {4, 1, 7}, {1, 65, 1}})
	test.Print()
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
				(*m).Max = value
			} else {
				if m.Max < value {
					(*m).Max = value
				}
			}
			(*m).Ancho = int64(j + 1)
		}
		(*m).Alto = int64(i + 1)
	}
	(*m).Matrix = matrix
	(*m).Cuadratica = (*m).Ancho == (*m).Alto
}

func (m *Matrix) Print() {

	for _, row := range (*m).Matrix {
		for j, value := range row {
			if j != int(m.Ancho-1) {
				fmt.Printf("%d ", value)
			} else {
				fmt.Printf("%d\n", value)
			}
		}
	}
}
