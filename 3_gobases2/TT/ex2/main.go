package main

import (
	"fmt"
)

type Matrix struct {
	content   [][]float64
	rows      int
	columns   int
	cuadratic bool
	maxValue  float64
}

func (m *Matrix) Set(content [][]float64) {
	m.content = content
	m.rows = len(content)
	m.columns = len(content[0])
	m.cuadratic = m.rows == m.columns
	m.maxValue = content[0][0]
	for _, row := range m.content {
		for _, e := range row {
			if e > m.maxValue {
				m.maxValue = e
			}
		}
	}
}

func (m Matrix) Print() {
	fmt.Print("Matriz de ", m.rows, "x", m.columns, " ")
	if m.cuadratic {
		fmt.Println("Cuadratica")
	} else {
		fmt.Println()
	}
	for _, row := range m.content {
		for _, e := range row {
			fmt.Print(e, "   ")
		}
		fmt.Println()
	}
	fmt.Println("Valor máximo: ", m.maxValue)
}

func main() {
	matrix1 := Matrix{}

	matrix1.Set([][]float64{
		{1, -2, 3.3},
		{4, -5.5, 6.6},
		{-7, 0.8, 9.99},
		{10, 11.11, -12.21},
	})
	matrix1.Print()
	fmt.Println()
	matrix2 := Matrix{}

	matrix2.Set([][]float64{
		{1, -2, 3.3},
		{4, -5.5, 6.6},
		{-7, 0.8, 9.99},
	})
	matrix2.Print()
}
