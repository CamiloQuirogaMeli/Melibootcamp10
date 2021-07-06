package main

import "fmt"

type Matrix struct {
	matriz     [][]float64
	filas      int
	columnas   int
	cuadratica bool
	maximo     float64
}

func (m *Matrix) set(matrixNueva [][]float64) {
	m.matriz = matrixNueva
	m.filas = len(matrixNueva[0])
	m.columnas = len(matrixNueva)
	m.cuadratica = (m.filas == m.columnas)
	m.maximo = 0
	for i := 0; i < len(matrixNueva); i++ {
		for j := 0; j < len(matrixNueva[i]); j++ {
			if m.maximo < matrixNueva[i][j] {
				m.maximo = matrixNueva[i][j]
			}
		}
	}

}

func (a Matrix) print() {
	for i := 0; i < len(a.matriz); i++ {
		for j := 0; j < len(a.matriz[i]); j++ {
			fmt.Print("|", a.matriz[i][j], "|")
		}
		fmt.Println("")
	}
	fmt.Println("Altura: ", a.columnas)
	fmt.Println("Ancho: ", a.filas)
	if a.cuadratica {
		fmt.Println("Es cuadratica")
	} else {
		fmt.Println("No es cuadratica")
	}
	fmt.Println("Valor maximo: ", a.maximo)
}
