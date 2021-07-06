package main

import "fmt"

/*
Una empresa de inteligencia artificial necesita tener una funcionalidad para crear matrices
en Go.
Para ello requieren una estructura Matrix que tenga los métodos:
- Set: Recibe una matrix de flotantes e inicializa los valores en la estructura Matrix
- Print: Imprime por pantalla la matrix de una formas más visible (Con los saltos de línea
entre filas)
La estructura matrix debe guardar la matrix, el tamaño de alto, el tamaño de ancho, si es
cuadrática y el valor máximo.
*/

const f = 2
const c = 2

type matrix struct {
	matriz      [f][c]int
	tamAlto     int
	tamAncho    int
	esCuadrada  bool
	valorMaximo int
}

func (m *matrix) set(datos [f][c]int) {

	m.matriz = datos
	m.tamAlto = f
	m.tamAncho = c
	m.esCuadrada = (f == c)
	m.valorMaximo = buscaMaximo(datos)

}

func buscaMaximo(datos [f][c]int) int {

	max := datos[0][0]
	for i := 0; i < f; i++ {
		for j := 0; j < c; j++ {

			if datos[i][j] > max {
				max = datos[i][j]
			}

		}
	}

	return max
}

func (m *matrix) print() {
	for i := 0; i < f; i++ {

		fmt.Println(m.matriz[i][:])

	}
}

func main() {

	var datos [f][c]int
	datos[0][0] = 1
	datos[0][1] = 2
	datos[1][0] = 3
	datos[1][1] = 4

	m1 := matrix{}

	m1.set(datos)
	fmt.Println(m1)

	m1.print()

}
