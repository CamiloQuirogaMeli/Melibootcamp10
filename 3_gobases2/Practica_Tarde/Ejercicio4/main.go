package main

import (
	"fmt"
	"reflect"
)

const (
	CHICO    = "chico"
	MEDIANO  = "mediano"
	GRANDE   = "grande"
	ESPECIAL = "especial"
	FRAGIL   = "fragil"
)

type Producto interface {
	tamTotal() float64
}

type ChicoProd struct {
	tam float64
}

type MedianoProd struct {
	tam float64
}

type GrandeProd struct {
	tam float64
}

type EspecialProd struct {
	tam float64
}

type FragilProd struct {
	tam float64
}

func (c ChicoProd) tamTotal() float64 {
	return c.tam
}
func (m MedianoProd) tamTotal() float64 {
	return m.tam + (m.tam * 0.05)
}
func (g GrandeProd) tamTotal() float64 {
	return g.tam + (g.tam * 0.20)
}
func (e EspecialProd) tamTotal() float64 {
	return e.tam + (e.tam * 0.75)
}
func (f FragilProd) tamTotal() float64 {
	return f.tam
}

type Flete struct {
	Productos []Producto
}

func (f *Flete) agregarProducto(prod ...Producto) {
	f.Productos = append(f.Productos, prod...)
}

//calcular la cant de envios sabiendo que puede llevar hasta 10.000.000 cm3
func (f *Flete) calcularEnvios(capacidad int) (cantEnviosEsp int, cantEnviosResto int) {
	var especial, resto []Producto
	var cantEsp, cantResto int
	for _, values := range f.Productos {
		if reflect.TypeOf(values).Name() == "EspecialProd" {
			especial = append(especial, values)
		} else {
			resto = append(resto, values)
		}
	}
	for _, values := range especial {
		cantEsp += int(values.tamTotal())
	}
	for _, values := range resto {
		cantResto += int(values.tamTotal())
	}
	cantEnviosEsp = cantEsp / capacidad
	if (cantEsp % capacidad) > 0 {
		cantEnviosEsp = cantEnviosEsp + 1
	}
	cantEnviosResto = cantResto / capacidad
	if (cantResto % capacidad) > 0 {
		cantEnviosResto = cantEnviosResto + 1
	}

	return
}

func factory(tam string, cm float64) Producto {
	switch tam {
	case CHICO:
		return ChicoProd{cm}
	case MEDIANO:
		return MedianoProd{cm}
	case GRANDE:
		return GrandeProd{cm}
	case FRAGIL:
		return FragilProd{cm}
	case ESPECIAL:
		return EspecialProd{cm}
	}
	return nil
}

func main() {
	p1 := factory(CHICO, 99)
	p2 := factory(MEDIANO, 500)
	p3 := factory(GRANDE, 1000)
	p4 := factory(FRAGIL, 200)
	p5 := factory(ESPECIAL, 1000)

	flete := Flete{}
	flete.agregarProducto(p1, p2, p3, p4, p5)

	EnvioEsp, EnvioResto := flete.calcularEnvios(100)
	fmt.Println("Cantidad de envios especiales:", EnvioEsp)
	fmt.Println("Cantidad de envios no especiales:", EnvioResto)

}
