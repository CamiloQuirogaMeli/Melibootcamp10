package main

import (
	"fmt"
)

type Matrix struct {
	matrix    [][]float64
	height    int
	weight    int
	quadratic bool
	max       float64
}

func (m *Matrix) setMatrix(theMatrix [][]float64) {
	m.matrix = theMatrix
	m.height = len(theMatrix)
	m.weight = len(theMatrix[0])
	m.quadratic = false // Inicializa la variable para posteriormente comprobar si es cuadratica o no
	for _, value := range m.matrix {
		for _, val := range value {
			if m.max < val {
				m.max = val
			}
		}
	}
	if m.height == m.weight {
		m.quadratic = true
	}
}

func details(m Matrix) {
	for _, vali := range m.matrix {
		fmt.Println(vali)
	}
	fmt.Println(m.height, m.weight, m.quadratic, m.max)
}

func main() {
	m := Matrix{}
	m.setMatrix([][]float64{{1, 21, 3, 4}, {4, 5, 6, 7}, {7, 8, 9, 9}, {4, 5, 68, 7}})
	details(m)
}
