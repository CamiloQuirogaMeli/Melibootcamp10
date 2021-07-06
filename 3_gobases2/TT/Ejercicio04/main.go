package main

import (
	"fmt"
	"math"
)

type Producto struct {
	Tamano float64
	Tipo   string
}

type Flete struct {
	Capacidad float64
	Productos [2]float64
}

func (p Producto) TamanoTotal() float64 {
	tamanoTotal := p.Tamano
	switch p.Tipo {
	case "Mediano":
		tamanoTotal = p.Tamano + (p.Tamano * 5 / 100)
	case "Grande":
		tamanoTotal = p.Tamano + (p.Tamano * 20 / 100)
	case "Fragil":
		tamanoTotal = p.Tamano + (p.Tamano * 75 / 100)
	}
	return tamanoTotal
}

func (f *Flete) AgregarProducto(p Producto) {
	if p.Tipo == "Especial" {
		f.Productos[0] += p.Tamano
	} else {
		f.Productos[1] += p.Tamano
	}
}

func (f Flete) CalcularEnvio() int {
	totalEnvios := 0
	for _, element := range f.Productos {
		totalEnvios += int(math.Ceil(element / f.Capacidad))
	}
	return totalEnvios
}

func main() {
	producto1 := Producto{1000, "Mediano"}
	producto2 := Producto{50000, "Grande"}
	producto3 := Producto{50000, "Especial"}
	flete := Flete{Capacidad: 10000}
	flete.AgregarProducto(producto1)
	flete.AgregarProducto(producto2)
	flete.AgregarProducto(producto3)
	fmt.Println("Total viajes", flete.CalcularEnvio())
}
