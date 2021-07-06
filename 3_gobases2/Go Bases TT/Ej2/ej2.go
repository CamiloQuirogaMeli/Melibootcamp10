package main

import "fmt"

type Matrix struct {
	matriz [][]float32
	alto   int
	ancho  int
}

func (m Matrix) Cuadratica() bool {
	return m.alto == m.ancho
}
func (m Matrix) Maximo() float32 {
	var maximo float32
	for _, fila := range m.matriz {
		for _, columna := range fila {
			if columna > maximo {
				maximo = columna
			}
		}
	}
	return maximo
}
func (m *Matrix) Set(val Matrix) {

	for i := 0; i < val.alto; i++ {
		fila := []float32{}
		for j := 0; j < val.ancho; j++ {
			var valor float32
			fmt.Printf("Valor %d:%d : ", i, j)
			fmt.Scanln(&valor)
			fila = append(fila, valor)
		}
		m.matriz = append(m.matriz, fila)
	}
}
func (m Matrix) Print() {
	for _, fila := range m.matriz {
		fmt.Printf("[")
		for key, valor := range fila {
			fmt.Printf("%.2f ", valor)
			if key+1 == m.ancho {
				fmt.Print("]\n")
			}
		}
	}
	fmt.Println("Maximo: ", m.Maximo())
}
func main() {
	variable := Matrix{}
	var alto, ancho int
	fmt.Println("Ancho: ")
	fmt.Scanln(&ancho)

	fmt.Println("Alto: ")
	fmt.Scanln(&alto)

	variable.alto = alto
	variable.ancho = ancho

	variable.Set(variable)
	variable.Print()
}
