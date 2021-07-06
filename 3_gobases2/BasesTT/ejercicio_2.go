package main

import (
	"fmt"
)

type Matriz struct {
	valores    [][]float64
	alto       int
	ancho      int
	maximo     float64
	esCuadrada bool
}

func (m *Matriz) Set(matriz [][]float64) {
	m.valores = matriz
	m.alto = len(matriz)
	m.ancho = len(matriz[0])
	m.esCuadrada = m.alto == m.ancho
	for _, fila := range matriz {
		for _, columna := range fila {
			if m.maximo < columna {
				m.maximo = columna
			}
		}
	}
}

func (m *Matriz) Print() {
	for _, fila := range m.valores {
		for _, columna := range fila {
			fmt.Printf("%v ", columna)
		}
		fmt.Printf("\n")
	}
}

func main() {
	m := Matriz{}
	m.Set([][]float64{[]float64{1, 2, 3, 1}, []float64{4, 5, 6, 0}, []float64{7, 8, 9, 4}})
	m.Print()
	fmt.Println("Ancho:", m.ancho)
	fmt.Println("Alto:", m.alto)
	fmt.Println("Maximo:", m.maximo)
	fmt.Println("Es cuadrada:", m.esCuadrada)
}
