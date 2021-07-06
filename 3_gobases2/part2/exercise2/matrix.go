package main

import (
	f "fmt"
)

type Matrix struct {
	height    int
	weight    int
	quadratic bool
	maxValue  float32
	data      [][]float32
}

func (matrix *Matrix) Set(data [][]float32) {
	matrix.height = len(data)
	matrix.weight = len(data[0])
	matrix.quadratic = matrix.height == matrix.weight
	matrix.data = data
	for _, row := range data {
		for _, element := range row {
			matrix.maxValue = row[0]
			if element > matrix.maxValue {
				matrix.maxValue = element
			}
		}
	}
}

func (matrix Matrix) Print() {
	for _, row := range matrix.data {
		f.Print("[")
		for i, element := range row {
			last := len(row)
			if last == i+1 {
				f.Print(element)
			} else {
				f.Print(element, ",")
			}
		}
		f.Println("]")
	}
}

func main() {
	var data [][]float32

	data = [][]float32{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	var matrix Matrix = Matrix{}
	matrix.Set(data)

	f.Println("Height:", matrix.height)
	f.Println("Weight:", matrix.weight)
	f.Println("Quadratic:", matrix.quadratic)
	f.Println("Max Value:", matrix.maxValue)
	f.Println("Matrix:")
	matrix.Print()
}
