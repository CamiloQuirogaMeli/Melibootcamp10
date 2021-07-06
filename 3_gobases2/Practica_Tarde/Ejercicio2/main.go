package main

import "fmt"

type Matrix struct {
	Matriz   [][]float64
	Alto     int
	Ancho    int
	IsCuadr  bool
	ValorMax float64
}

func (m *Matrix) SetMatrix(valores [][]float64) {
	m.Matriz = valores
	m.Alto = len(valores)
	m.Ancho = len(valores[0])
	m.ValorMax = valorMax(valores)
	m.IsCuadr = true
}

func valorMax(valores [][]float64) (max float64) {
	for i := range valores {
		for j := range valores[0] {
			if valores[i][j] > max {
				max = valores[i][j]
			}
		}
	}
	return
}

func (m Matrix) print() {
	for i := range m.Matriz {
		for j := range m.Matriz[0] {
			fmt.Printf("%.2f ", m.Matriz[i][j])
		}
		fmt.Print("\n\n")
	}
}

func main() {
	var a = [][]float64{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}}
	m1 := Matrix{}
	m1.SetMatrix(a)
	m1.print()

	fmt.Print("---------------------\n")
	fmt.Print("Matriz: ", m1.Matriz, "\n")
	fmt.Print("Alto: ", m1.Alto, "\n")
	fmt.Print("Ancho: ", m1.Ancho, "\n")
	fmt.Print("Cuadratica: ", m1.IsCuadr, "\n")
}
