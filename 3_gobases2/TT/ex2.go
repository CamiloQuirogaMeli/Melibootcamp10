package main

import (
	"fmt"
)

type Matrix struct {
	Values      [][]float64
	Heigth      int
	Width       int
	IsQuadratic bool
	MaxValue    float64
}

func (matrix *Matrix) set(data [][]float64) {
	matrix.Values = data
	matrix.Heigth = len(data)
	matrix.Width = len(data[0])
	matrix.IsQuadratic = matrix.Heigth == matrix.Width

	matrix.MaxValue = matrix.Values[0][0]
	for _, row := range matrix.Values {
		for _, element := range row {
			if element > matrix.MaxValue {
				matrix.MaxValue = element
			}
		}
	}
}

func (matrix Matrix) print() {
	fmt.Println("Alto:", matrix.Heigth)
	fmt.Println("Ancho:", matrix.Width)
	fmt.Println("Cuadratica?:", matrix.IsQuadratic)
	fmt.Println("Valor máximo:", matrix.MaxValue)

	for _, row := range matrix.Values {
		for _, element := range row {
			fmt.Printf("%.1f ", element)
		}
		fmt.Println()
	}
}

func main() {
	data := [][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{1, 1, 1, 1},
	}
	matrix := Matrix{}
	matrix.set(data)
	matrix.print()
}
