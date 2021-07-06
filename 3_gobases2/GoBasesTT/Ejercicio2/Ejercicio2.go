package main

import "fmt"

type matrix struct {
	valores     [100][100]float32
	filas       int
	columnas    int
	cuadratica  bool
	valorMaximo float32
}

func (m *matrix) Set(mat [100][100]float32, f int, c int) {
	may := mat[0][0]

	//Asignamos valores a la matriz y vemos cual es el mayor
	for i := 0; i < f; i++ {
		for j := 0; j < c; j++ {
			m.valores[i][j] = mat[i][j]
			if mat[i][j] > may {
				may = mat[i][j]
			}
		}
	}

	m.filas = f
	m.columnas = c
	m.cuadratica = esCuadratica(f, c)
	m.valorMaximo = may
}

func (m matrix) Print() {
	for i := 0; i < m.filas; i++ {
		for j := 0; j < m.columnas; j++ {
			fmt.Printf("%.1f ", m.valores[i][j])
		}
		fmt.Printf("\n")
	}
}

func esCuadratica(f int, c int) bool {
	if f == c {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println("Inicio")

	matriz := [100][100]float32{{1.1, 2.2, 3.3}, {4.4, 5.5, 6.6}, {7.7, 8.8, 9.9}}

	ma := matrix{}
	ma.Set(matriz, 3, 3)
	ma.Print()
}
