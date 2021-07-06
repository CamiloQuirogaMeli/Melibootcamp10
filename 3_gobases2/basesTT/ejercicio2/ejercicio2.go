package main

import "fmt"

type Matrix struct {
	matriz     [][]float64
	altura     int
	ancho      int
	cuadratica bool
	maximo     float64
}

func (m *Matrix) SetMatrix(matrixNueva [][]float64) {
	m.matriz = matrixNueva
	m.altura = len(matrixNueva)
	m.ancho = len(matrixNueva[0])
	m.cuadratica = (m.altura == m.ancho)
	m.maximo = 0
	for i := 0; i < len(matrixNueva); i++ {
		for j := 0; j < len(matrixNueva[i]); j++ {
			if m.maximo < matrixNueva[i][j] {
				m.maximo = matrixNueva[i][j]
			}
		}
	}
}

func (m Matrix) PrintMatrix() {
	for i := 0; i < len(m.matriz); i++ {
		for j := 0; j < len(m.matriz[i]); j++ {
			fmt.Print("|", m.matriz[i][j], "|")
		}
		fmt.Println("")
	}
	fmt.Println("Altura: ", m.altura)
	fmt.Println("Ancho: ", m.ancho)
	if m.cuadratica {
		fmt.Println("Es cuadratica")
	} else {
		fmt.Println("No es cuadratica")
	}
	fmt.Println("Valor maximo: ", m.maximo)
}

func main() {
	a := []float64{1, 2, 3}
	b := []float64{4, 5, 6}
	matrix := [][]float64{a, b}
	var m Matrix
	m.SetMatrix(matrix)
	m.PrintMatrix()

}
