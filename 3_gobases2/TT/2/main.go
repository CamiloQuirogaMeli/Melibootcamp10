package main

import (
	"fmt"
)

func main() {

	m := Matrix{}

	m.Set([][]float64{{1, 2, 3}, {4, 5, 6}, {6, 7, 1}})
	m.Print()
}

type Matrix struct {
	Values    [][]float64
	Height    int
	Width     int
	Quadratic bool
	MaxValue  float64
}

func (m *Matrix) Set(values [][]float64) {
	m.Height = len(values)
	m.Width = len(values[0])
	m.Quadratic = m.Height == 3 && m.Width == 3
	m.Values = values
	m.MaxValue = FindMax(values)
}

func (m Matrix) Print() {
	for _, i := range m.Values {
		for _, j := range i {
			fmt.Print(j, "\t")
		}
		fmt.Print("\n")
	}
	fmt.Println("El valor maximo de la matriz es ", m.MaxValue)
	fmt.Printf("El alto de la matriz es %d y el anch es %d \n", m.Height, m.Width)
	if m.Quadratic {
		fmt.Println("La matrix es cuadratica")
	}
}

func FindMax(values [][]float64) float64 {
	max := values[0][0]

	for _, i := range values {
		for _, j := range i {
			if j > max {
				max = j
			}
		}
	}

	return max
}
