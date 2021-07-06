package main

import "fmt"

type Matrix struct {
	Matriz      [][]float64
	Alto        int
	Ancho       int
	Cuadratica  bool
	ValorMaximo float64
}

func (m *Matrix) Set(matriz [][]float64) {
	m.Matriz = matriz
	m.Alto = len(matriz)
	m.Ancho = len(matriz[0])
	m.Cuadratica = len(matriz[0]) == len(matriz)
	m.ValorMaximo = encontrarMaximo(matriz)
}

func (m Matrix) Print() {
	for i := 0; i < m.Alto; i++ {
		for j := 0; j < m.Ancho; j++ {
			fmt.Printf("%.2f\t", m.Matriz[i][j])
		}
		fmt.Printf("\n")
	}
}

func encontrarMaximo(matriz [][]float64) float64 {
	max := matriz[0][0]
	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[0]); j++ {
			if matriz[i][j] > max {
				max = matriz[i][j]
			}
		}
	}
	return max
}

func main() {
	var m1 Matrix
	unaMatriz := [][]float64{
		{1, 8, 4, 5},
		{45, -6, 7, 15},
		{1, 2, 0, 33},
	}
	var m2 Matrix
	otraMatriz := [][]float64{
		{0, 1},
		{7, 2},
	}
	m1.Set(unaMatriz)
	m2.Set(otraMatriz)

	m1.Print()
	fmt.Println("Cuadratica: ", m1.Cuadratica)
	fmt.Println("Alto: ", m1.Alto)
	fmt.Println("Ancho: ", m1.Ancho)
	fmt.Println("Valor máximo: ", m1.ValorMaximo)
	fmt.Printf("\n")

	m2.Print()
	fmt.Println("Cuadratica: ", m2.Cuadratica)
	fmt.Println("Alto: ", m2.Alto)
	fmt.Println("Ancho: ", m2.Ancho)
	fmt.Println("Valor máximo: ", m2.ValorMaximo)
}
