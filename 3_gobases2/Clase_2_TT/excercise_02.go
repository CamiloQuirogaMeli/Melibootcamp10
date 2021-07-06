package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	mat, err := SetMatrix([]float64{2.23, 1.15, 2.33, 3.52}, 2, 2)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	mat.printMatrix()
	fmt.Printf("Max Value: [%2.f]\n", mat.max)
	fmt.Printf("Matrix isCuadratic: [%v]\n", mat.isCuadratic)
}

type Matrix struct {
	matriz      []float64
	width       int
	height      int
	isCuadratic bool
	max         float64
}

func SetMatrix(matriz []float64, width, height int) (Matrix, error) {
	m := Matrix{}
	if len(matriz) == 0 {
		return Matrix{}, errors.New("matriz cannot be empty")
	}

	if width <= 0 || height <= 0 {
		return Matrix{}, errors.New("Width and height cannot be negative or zero")
	}

	m.matriz = matriz
	m.width = width
	m.height = height
	m.isCuadratic = isCuadratic(m.width, m.height)
	m.max = m.getMax()

	return m, nil
}

func (m Matrix) printMatrix() {
	if len(m.matriz) == 0 {
		fmt.Println("La matriz está vacía")
	} else {
		for width := 0; width < m.width; width++ {
			fmt.Println(m.matriz[width*m.height : width*m.height+m.height])
		}
	}
}

func isCuadratic(width, height int) bool {
	if width != height {
		return false
	}
	return true
}

func (m Matrix) getMax() float64 {
	var maxValue float64
	for i, value := range m.matriz {
		if i == 0 || value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}
