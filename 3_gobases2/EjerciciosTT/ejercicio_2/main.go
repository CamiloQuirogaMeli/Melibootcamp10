package main

import "fmt"

type Matrix struct {
	filas        int
	columnas     int
	esCuadratica bool
	valorMaximo  float64
}

func (m *Matrix) set(matrix [5][2]float64) {
	filas := len(matrix)
	columnas := len(matrix[0])

	m.filas = filas
	m.columnas = columnas

	m.esCuadratica = false

	var i, j int
	var maximo float64 = matrix[0][0]

	/* output each array element's value */
	for i = 1; i < filas; i++ {
		for j = 0; j < columnas; j++ {
			if matrix[i][j] > maximo {
				maximo = matrix[i][j]
			}
		}
	}

	m.valorMaximo = maximo
}

func (m Matrix) print() {
	fmt.Println("Detalle de la Matriz: ")
	fmt.Println("Filas: ", m.filas)
	fmt.Println("Columnas: ", m.columnas)
	fmt.Println("Es Cuadratica: ", m.esCuadratica)
	fmt.Println("Valor Maximo: ", m.valorMaximo)
}

func main() {
	var m = [5][2]float64{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}

	m2 := Matrix{}

	m2.set(m)
	m2.print()
}
