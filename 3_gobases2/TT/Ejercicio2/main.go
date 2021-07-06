package main

import "fmt"

type Matrix struct {
	Matriz   [][]float64
	Alto     int
	Ancho    int
	ValorMax float64
}

func (m *Matrix) setMatrix(valores [][]float64) {
	m.Matriz = valores
	m.Alto = len(valores)
	m.Ancho = len(valores[0])
	m.ValorMax = setearValorMAx(valores)
}

func (m Matrix) printMatrix() {
	fmt.Print("\nMatrix is: \n\n")
	for i := 0; i < len(m.Matriz); i++ {
		for j := 0; j < len(m.Matriz[0]); j++ {
			fmt.Printf("%.2f ", m.Matriz[i][j])
		}
		fmt.Print("\n\n")
	}

	if m.Alto == m.Ancho {
		fmt.Println("La matriz es cuadrática")
	}
}

func setearValorMAx(valores [][]float64) (valMax float64) {
	var max float64 = 0
	var j, i int

	for i = 0; i < len(valores); i++ {
		for j = 0; j < len(valores[0]); j++ {
			if valores[i][j] > max {
				max = valores[i][j]
			}
		}
	}
	valMax = max

	return
}

func main() {
	var a = [][]float64{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	m1 := Matrix{}
	m1.setMatrix(a)

	m2 := Matrix{}
	m2.setMatrix([][]float64{{1.5, 1.9, 9.2}, {7.8, 7.6, 0.5}, {4.8, 2.2, 3.5}})

	fmt.Print(m1)
	m1.printMatrix()

	fmt.Println("--------------")

	fmt.Print(m2)
	m2.printMatrix()
}
