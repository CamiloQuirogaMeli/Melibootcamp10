package main

import (
	f "fmt"
)

type Matrix struct {
	Values    []float64
	Rows      int
	Columns   int
	Quadratic bool
	Max_value float64
}

func main() {
	matrix := Matrix{}

	f.Print("Enter the number of rows: ")
	f.Scanln(&matrix.Rows)

	f.Print("Enter the number of columns: ")
	f.Scanln(&matrix.Columns)

	f.Println("Enter the matrix values ")
	matrix.setValues()

	matrix.determineQuadratic()
	matrix.findMax_value()

	matrix.printMatrix()
	matrix.getFeatures()
}

func (m Matrix) getFeatures() {
	f.Printf("\n\nShape: (%d , %d)\n", m.Rows, m.Columns)
	f.Println("Quadratic: ", m.Quadratic)
	f.Println("Maximum value: ", m.Max_value)
}

func (m *Matrix) determineQuadratic() {
	if m.Rows == m.Columns {
		m.Quadratic = true
	} else {
		m.Quadratic = false
	}
}

func (m *Matrix) findMax_value() {
	var max float64 = m.Values[0]

	for _, value := range m.Values {
		if value > max {
			max = value
		}
	}
	m.Max_value = max
}

func (m *Matrix) setValues() {
	var a float64
	k := 0
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Columns; j++ {
			f.Printf("[%d, %d]: ", i, j)
			f.Scanln(&a)
			m.Values = append(m.Values, a)
			k++
		}
	}
}

func (m Matrix) printMatrix() {
	k := 0
	for i := 0; i < m.Rows; i++ {
		print("\n")
		for j := 0; j < m.Columns; j++ {
			f.Print(m.Values[k], " ")
			k++
		}
	}
}
