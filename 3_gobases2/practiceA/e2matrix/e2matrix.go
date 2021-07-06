package main

import (
	"fmt"
)

type Matrix struct {
	Values     [][]float64
	Alto       int
	Ancho      int
	Cuadratica bool
	Maximo     float64
}

func main() {
	m := Matrix{}
	aux := [][]float64{{11, 12, 13, 14}, {21, 22, 23, 24}, {31, 32, 33, 34}}
	m.set(aux)
	m.print()
}

func (m *Matrix) set(values [][]float64) {
	m.Values = make([][]float64, len(values))
	for key, value := range values {
		m.Values[key] = make([]float64, len(value))
		copy(m.Values[key], value)
	}
	m.Alto = len(values)
	m.Ancho = len(values[0])
	m.Cuadratica = false // Más tarde armar una rutina para comprobralo

	var max float64
	flag := true
	for _, rows := range values {
		for _, column := range rows {
			if flag {
				max = column
				flag = false
			} else {
				if column > max {
					max = column
				}
			}
		}
	}
	m.Maximo = max

}

func (m Matrix) print() {
	fmt.Println("Cuadrática: ", m.Cuadratica)
	fmt.Println("Máximo: ", m.Maximo)
	fmt.Println("Ancho: ", m.Ancho)
	fmt.Println("Alto: ", m.Alto)

	for _, rows := range m.Values {
		fmt.Println(rows)
	}

}
