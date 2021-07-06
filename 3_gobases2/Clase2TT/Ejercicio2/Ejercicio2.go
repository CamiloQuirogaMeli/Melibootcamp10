package main

import "fmt"

//Una empresa de inteligencia artificial necesita tener una funcionalidad para crear matrices en Go. Para ello requieren una estructura Matrix que tenga los métodos:
//- Set: Recibe una matrix de flotantes e inicializa los valores en la estructura Matrix
//- Print: Imprime por pantalla la matrix de una formas más visible (Con los saltos de línea entre filas)
//La estructura matrix debe guardar la matrix, el tamaño de alto, el tamaño de ancho, si es cuadrática y el valor máximo.

type Matrix struct {
	matrix     string
	Alto       float64
	Ancho      float64
	Cuadratica bool
	ValMax     float64
}

func (m Matrix) detalle() {
	fmt.Println("Matrix: [", m.matrix, "]\nAlto: [", m.Alto, "]\nAncho: [", m.Ancho, "]\nCuadratica: [", m.Cuadratica, "]\nValor Maximo: [", m.ValMax, "]")
}

func main() {

	m1 := Matrix{"m1", 4, 4, true, 8}

	m1.detalle()

}
