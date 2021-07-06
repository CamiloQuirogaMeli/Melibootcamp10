package main

import "fmt"

type Matrix struct {
	values [][]int
}

func (m Matrix) height() int {
	var height int
	height = len(m.values)
	return height
}

func (m Matrix) width() int {
	var width int
	width = len(m.values[0])
	return width
}

func (m Matrix) isQuadratic() bool {
	if m.height() == m.width() {
		return true
	}
	return false
}

func (m Matrix) maxValue() int {
	var max int
	for i := 0; i < len(m.values); i++ {
		for j := 0; j < len(m.values[i]); j++ {
			if m.values[i][j] > max {
				max = m.values[i][j]
			}
		}
	}
	return max
}

func (m *Matrix) setValues(values [][]int) {
	m.values = values
}

func (m Matrix) Print() {
	for i := 0; i < len(m.values); i++ {
		for j := 0; j < len(m.values[i]); j++ {
			fmt.Print(m.values[i][j], " ")
		}
		fmt.Println()
	}
}

func main() {

	m := Matrix{}
	data := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	m.setValues(data)

	fmt.Println("Alto", m.height())
	fmt.Println("Ancho", m.width())
	fmt.Println("Es Cuadratica: ", m.isQuadratic())
	fmt.Println("Valor Maximo:", m.maxValue())
	m.Print()

}
