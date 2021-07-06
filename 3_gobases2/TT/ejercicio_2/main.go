package main

import (
	"fmt"
	"math"
)

type Matrix struct {
	Data [][]float64
	Height int
	Width int
	Quadratic bool
	Maximum float64
}

func (matrix *Matrix) set(data [][]float64) {
	matrix.Data = data
	matrix.Height = len(data)
	matrix.Width = len(data[0])
	matrix.Quadratic = (matrix.Height == matrix.Width)

	matrix.Maximum = matrix.Data[0][0]
	for _, row := range matrix.Data {
		for _, elem := range row {
			matrix.Maximum = math.Max(matrix.Maximum, elem)
		}
	}
}

func (matrix Matrix) print() {

	for _, row := range matrix.Data {
		fmt.Print("[ ")
		for _, elem := range row {
			fmt.Printf("%.2f ", elem)
		}
		fmt.Println("]")
	}
}

func main() {

	rows := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	var aMatrix Matrix = Matrix{}
	aMatrix.set(rows)

	fmt.Println("Alto:", aMatrix.Height)
	fmt.Println("Ancho:", aMatrix.Width)
	fmt.Println("Cuadratica:", aMatrix.Quadratic)
	fmt.Println("Maximo:", aMatrix.Maximum)
	fmt.Println("Contenido:")
	aMatrix.print()
}