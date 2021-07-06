package main

import "fmt"

const (
	chico   = "chico"
	mediano = "mediano"
	grande  = "grande"
)

type Producto interface {
	precio() float64
}

type Productos struct {
	costo float64
}

func (p Productos) precio() float64 {
	return p.costo
}

func factory(tipoProducto string, precio float64) float64 {
	switch tipoProducto {
	case chico:
		return precio
	case mediano:
		return precio + (precio * 0.3)
	case grande:
		return precio + (precio * 0.6) + 2500
	default:
		return 0
	}
}

func mostrarPrecio(p Producto) {
	fmt.Println(p)
	fmt.Println(p.precio())
}

func main() {
	producto := Productos{costo: 150}

	mostrarPrecio(producto)
}
