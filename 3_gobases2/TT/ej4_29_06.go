package main

import (
	"errors"
	"fmt"
)

type Flete struct {
	capacidadOcupada  float64
	capacidadMaximo   float64
	cantidadProductos int
}

type Producto struct {
	Tamaño string
	Costo  float64
	C3     float64
}

func (p Producto) TamanoTotal() float64 {
	if p.Tamaño == "chico" {
		return p.C3
	} else if p.Tamaño == "mediano" {
		return p.C3 * 1.05
	} else if p.Tamaño == "grande" {
		return p.C3 * 1.2
	} else if p.Tamaño == "fragil" {
		return p.C3 * 1.75
	} else if p.Tamaño == "especial" {
		return 0
	}
	return 0
}

func (f Flete) AgregarProducto(p Producto) (float64, error) {
	if (f.capacidadOcupada + p.C3) < 10000000 {
		var cap = f.capacidadOcupada + p.C3
		f.cantidadProductos++
		return cap, nil
	} else {
		return f.capacidadOcupada, errors.New("la capacidad ha sido agotada")
	}
}

func (f Flete) CalcularEnvios() int {
	return (f.cantidadProductos)
}

func main() {
	p := Producto{"grande", 100, 11000}
	f := Flete{0, 1000, 0}
	var pre, err = f.AgregarProducto(p)
	fmt.Println("La capacidad ocupada del flete es de: ", pre, "cm3")
	fmt.Println(err)
	var tot = p.TamanoTotal()
	fmt.Println("El tamaño total del producto es: ", tot)
	var tot1 = f.CalcularEnvios()
	fmt.Println("la cantidad de envios son", tot1)
}
