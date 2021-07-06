package main

import "fmt"

type Matrix struct {
	Valores    [][]float64
	Alto       int
	Ancho      int
	Cuadratica bool
	Maximo     float64
}

func (m *Matrix) set(values [][]float64, h int, w int) {
	var max float64
	m.Valores = values
	m.Alto = h
	m.Ancho = w
	m.Cuadratica = (h == w)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if i == 0 && j == 0 {
				max = values[i][j]
			} else {
				if values[i][j] > max {
					max = values[i][j]
				}
			}
		}

	}
	m.Maximo = max
}

func (m Matrix) detalle() {
	for i := 0; i < m.Alto; i++ {
		for j := 0; j < m.Ancho; j++ {
			fmt.Printf("%.1f\t", m.Valores[i][j])
		}
		fmt.Println()
	}
}

func main() {
	var a = [][]float64{{0.2, 0.3, 1.0}, {1.6, 2.4, 2.7}, {2.0, 4.3, 4.4}, {3.0, 6.0, 5.1}, {4.1, 8.5, 9.9}}
	m := Matrix{}
	m.set(a, 5, 3)
	m.detalle()
}
