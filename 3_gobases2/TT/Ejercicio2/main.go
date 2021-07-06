package main

import "fmt"

func main() {
	miMatriz := [][]float64{{2, 4, 6}, {1, 3, 5}, {8, 10, 12}} //es un slice multidimensional
	m := matrix{valores: miMatriz}
	m.set(miMatriz)

	fmt.Println(m.printMatrix())
}

type matrix struct {
	valores     [][]float64 `bd:"valores_matriz"`
	alto        int         `bd:"tamaño_altura"`
	ancho       int         `bd:"tamaño_ancho"`
	cuadratica  bool        `bd:"cuadratica"`
	valorMaximo float64     `bd:"valor_maximo"`
}

func (m *matrix) set(matriz [][]float64) { //al usar el puntero modifico la matrix de main
	var valorMaximo float64
	for i, val := range m.valores {
		for _, valj := range matriz[i] {
			if i != 0 {
				valorMaximo = valj
			} else {
				valorMaximo = val[0]
			}
		}
	}
	//	if condicionParaQueSeaCuadratica {
	//		m.cuadratica = true
	//	}

	m.alto = len(matriz)
	m.ancho = cap(matriz)
	m.cuadratica = false //lo hardeo para no perder tiempo con eso
	m.valorMaximo = valorMaximo

	//fmt.Println(m.printMatrix()) //esto tenia que hacerlo cuando no utilizaba el puntero
}

func (m matrix) printMatrix() ([][]float64, int, int, bool, float64) {
	return m.valores, m.alto, m.ancho, m.cuadratica, m.valorMaximo
}
