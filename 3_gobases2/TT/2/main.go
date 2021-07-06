/* Ejercicio 2 - Matrix
Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una
estructura que represente una matriz de datos.
Para ello requieren una estructura Matrix que tenga los métodos:
- Set: Recibe una matriz de flotantes e inicializa los valores en la estructura Matrix
- Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea
entre filas)
La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la
dimensión del ancho, si es cuadrática y cuál es el valor máximo. */
package main

import (
	"fmt"
)

type Matrix struct {
	height []float64
	width []float64
	cuadratic bool
	maxValue int
}
func (m *Matrix) Set(matriz [][]float64) {
	m.height = matriz[0]
	m.width = matriz[1]
	m.cuadratic = boolean(len(matriz[0]) == len(matriz[1]))
}
func (m Matrix) Print() {
	fmt.Printf("--- Matrix ---\nno se como hacer aca...")
}