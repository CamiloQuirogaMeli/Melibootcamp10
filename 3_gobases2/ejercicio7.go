package main

import (
	"fmt"
)

func ejercicio7() {
	elementos := []float64{99, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	newMatrix := Matrix{
		height: 5,
		width:  3,
	}

	newMatrix.set(elementos)
	newMatrix.print()
	fmt.Printf("\nEs cuadratica: %t\n", newMatrix.isCudratic())
	fmt.Printf("Es valor maximo es: %v\n", newMatrix.maxValue())

}

type Matrix struct {
	elementos     []float64
	height, width int
}

func (m *Matrix) set(matriz []float64) {
	m.elementos = matriz
}

func (m *Matrix) print() {
	var currentHeight int = 1
	var currentWidth int = 0
	for i := 0; i < len(m.elementos); i++ {
		if currentWidth < m.width {
			fmt.Printf("%v ", m.elementos[i])
			currentWidth++
		} else if currentWidth == m.width {
			fmt.Printf("\n%v ", m.elementos[i])
			currentHeight++
			currentWidth = 1
		} else if currentHeight == m.height {
			break
		}
	}
}

func (m *Matrix) isCudratic() bool {
	if m.width == m.height {
		return true
	} else {
		return false
	}
}

func (m *Matrix) maxValue() (maxElement float64) {
	maxElement = m.elementos[0]
	for _, elemento := range m.elementos {
		if elemento > maxElement {
			maxElement = elemento
		}
	}
	return
}
