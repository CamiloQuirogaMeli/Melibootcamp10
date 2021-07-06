package main

import (
	"fmt"
)

type Matrix struct {
	Alto        float32
	Ancho       float32
	cuadratica  bool
	valorMaximo float32
}

func (matrix *Matrix) setAlto(alto float32) {
	matrix.Alto = alto
}

func (matrix *Matrix) print() {

	fmt.Println(matrix)
	/*
		for i := 0; i < matrixFilas.NumField(); i++ {
			field := matrixFilas.Field(i)
			fmt.Println(field)
		}
	*/
}

func main() {

	miMatrix := Matrix{34.2, 32.1, false, 450.0}
	miMatrix.print()
	miMatrix.setAlto(50.0)
	miMatrix.print()

}
