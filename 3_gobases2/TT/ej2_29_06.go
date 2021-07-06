package main

import "fmt"

type Matrix struct {
	Alto       float64
	Ancho      float64
	Cuadratica bool
	Maximo     float64
}

func (a Matrix) Print() {
	fmt.Println("Alto", a.Alto)
	fmt.Println("Ancho", a.Ancho)
	if a.Alto == a.Ancho {
		fmt.Println("Cuadratica", a.Cuadratica)
	}
	fmt.Println("Maximo", a.Maximo)
}

func (m *Matrix) Set(al float64, an float64, cu bool, ma float64) {
	m.Alto = al
	m.Ancho = an
	m.Cuadratica = cu
	m.Maximo = ma
}

func main() {
	m := Matrix{}
	m.Set(10.0, 10.0, true, 40)

	m.Print()

}
