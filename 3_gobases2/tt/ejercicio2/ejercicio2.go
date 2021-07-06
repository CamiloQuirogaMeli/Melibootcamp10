package main

import (
	"fmt"
	"math/rand"
)

type Matrix struct {
	Alto   int
	Ancho  int
	Matriz []float32
}

func (m Matrix) Cuadratica() bool {
	if m.Alto == m.Ancho {
		return true
	} else {
		return false
	}
}

func (m Matrix) Max() float32 {
	var maximo float32 = 0
	for i := range m.Matriz {
		if m.Matriz[i] > maximo {
			maximo = m.Matriz[i]
		}
	}
	return maximo
}

func Set(m Matrix) Matrix {
	min := 1
	max := 30
	m.Matriz = make([]float32, m.Alto*m.Ancho)
	for i := 0; i < m.Ancho*m.Alto; i++ {
		m.Matriz[i] = float32(rand.Intn(max-min+1) + min)
	}
	return m
}

func Print(m Matrix) {
	for i := range m.Matriz {
		fmt.Printf(" %.2f ", m.Matriz[i])
		if i+1 == m.Ancho || (i+1)%m.Ancho == 0 {
			fmt.Printf("\n")
		}
	}
}

func main() {
	variable := Matrix{}
	variable.Alto = 8
	variable.Ancho = 8
	variable = Set(variable)
	Print(variable)

	cuadrada := variable.Cuadratica()
	if cuadrada {
		fmt.Printf(" \n La matriz es cuadrada")
	} else {
		fmt.Printf(" \n La matriz no es cuadrada")
	}

	fmt.Printf(" \n El valor máximo es %.2f\n", variable.Max())

}
