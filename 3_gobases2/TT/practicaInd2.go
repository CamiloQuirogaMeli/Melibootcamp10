package main

import (
	"fmt"
)

/* estructura Matrix que tenga los métodos:
Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho, si es cuadrática y cuál es el valor máximo.
*/

type Matrix struct {
	Valores    [][]float64
	Alto       int
	Ancho      int
	Cuadratica bool
	ValorMax   float64
}

func (m *Matrix) setDatos(valores [][]float64, alto int, ancho int) {

	m.Valores = valores
	m.Alto = alto
	m.Ancho = ancho
	m.Cuadratica = (alto == ancho)

	for i := 0; i < m.Alto; i++ {
		for j := 0; j < m.Ancho; j++ {
			if i == 0 && j == 0 {

				m.ValorMax = valores[i][j]
			} else {
				if valores[i][j] > m.ValorMax {
					m.ValorMax = valores[i][j]
				}
			}
		}

	}

}

func (m *Matrix) printDatos() {

	for i := 0; i < m.Alto; i++ {
		for j := 0; j < m.Ancho; j++ {
			fmt.Printf(" El valor de la FILA %v COLUMNA %v es: %.2f \n", i+1, j+1, m.Valores[i][j])
		}
	}

	fmt.Printf("Es una matriz de %v por %v \n", m.Alto, m.Ancho)

}

func main() {

	var valores = [][]float64{{3.9, 2.7, 3.1, 8.2}, {7.3, 9.1, 7.2, 3.9}, {7.1, 2.3, 1.2, 3.4}}
	m := Matrix{}

	m.setDatos(valores, 3, 4)

	m.printDatos()
}
