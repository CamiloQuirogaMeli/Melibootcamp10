package main

import "fmt"

type Matrix struct {
	value     [][]float64
	height    int
	width     int
	quadratic bool
	valMax    float64
}

func (m *Matrix) setMatrix(value [][]float64) {

	m.value = value
	m.height = len(m.value)
	m.valMax = value[0][0]

	for i := 0; i < m.height; i++ {
		m.width = len(m.value[i])
		for j := 0; j < m.width; j++ {
			var actual float64 = m.value[i][j]
			if actual > m.valMax {
				m.valMax = actual
			}
		}
	}
	m.quadratic = false
	if m.height == m.width {
		m.quadratic = true
	}
}

func (m Matrix) print() [][]float64 {
	for i := 0; i < m.height; i++ {
		for j := 0; j < m.width; j++ {
			fmt.Print(m.value[i][j], "\t")
		}
		fmt.Println()
	}
	return m.value
}

func main() {
	m1 := Matrix{}
	matriz := [][]float64{{1.3, 2.3, 3.6}, {4.4, 5.4, 6.5}, {7.8, 8.8, 9.2}}

	m1.setMatrix(matriz)
	m1.print()
}
