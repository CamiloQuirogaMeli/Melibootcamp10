package main

import (
	"fmt"
	"math"
)

const (
	CARGA_MAXIMA = 10000000
	CHICO        = "chico"
	MEDIANO      = "mediano"
	GRANDE       = "grande"
	FRAGIL       = "fragil"
	ESPECIAL     = "especial"
)

type Producto struct {
	tamanio   float64
	adicional float64
}

func (p Producto) TamanioTotal() float64 {
	return p.tamanio
}

type Flete struct {
	producto Producto
}

func (f *Flete) AgregarProducto(producto Producto) {
	f.producto = producto
}

func (f *Flete) CalcularEnvios() float64 {
	return math.Ceil(CARGA_MAXIMA / (f.producto.tamanio * (1 + f.producto.adicional)))
}

var productos = map[string]Producto{
	CHICO:    {1000.0, 0.0},
	MEDIANO:  {10000.0, 0.05},
	GRANDE:   {100000.0, 0.20},
	FRAGIL:   {5000.0, 0.75},
	ESPECIAL: {50000.0, 0.0},
}

func main() {
	p := productos[FRAGIL]
	flete := Flete{}
	flete.AgregarProducto(p)
	fmt.Println("Tamaño:", p.TamanioTotal())
	fmt.Println("Envíos:", flete.CalcularEnvios())
}
