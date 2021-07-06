package main

import (
	"errors"
	"fmt"
	"os"
)

type matrix struct {
	values      [][]float64
	dimCol      int
	dimRow      int
	isQuadratic bool
	maxValue    float64
}

func (m *matrix) Set(inputValues []float64, dimCol, dimRow int) error {
	if dimCol*dimRow != len(inputValues) {
		return errors.New("the matrix can't be filled")
	}
	m.dimCol = dimCol
	m.dimRow = dimRow
	if dimCol == dimRow {
		m.isQuadratic = true
	}
	// m.values = [dimRow][dimCol]float64{}
	index := 0
	max := inputValues[index]
	for i := 0; i < dimRow; i++ {
		m.values[i] = make([]float64, dimCol)
		for j := 0; j < dimCol; j++ {
			m.values[i][j] = inputValues[index]
			if inputValues[index] > max {
				max = inputValues[index]
			}
			index++
		}
	}
	m.maxValue = max
	return nil
}

func (m *matrix) Print() {
	fmt.Println("Matrix:")
	for _, v := range m.values {
		fmt.Println(v)
	}
	fmt.Println("Column dimension: ", m.dimCol)
	fmt.Println("Row dimension: ", m.dimRow)
	fmt.Println("Is quadratric? ", m.isQuadratic)
	fmt.Println("Max value: ", m.maxValue)
}

func main() {
	input := []float64{2, 5, 4, 6}
	var dimRow int = 2
	var dimCol int = 2
	var matrix = matrix{
		values: make([][]float64, dimCol),
	}
	err := matrix.Set(input, dimRow, dimCol)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	matrix.Print()
}
