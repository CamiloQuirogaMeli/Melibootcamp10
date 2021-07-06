package main

import (
	"reflect"
)

// Un Ecommerce necesita realizar una funcionalidad en Go para gestionar el envío y reparto de
// productos:
// La empresa tiene 5 tipos de productos: Chico, Mediano, Grande, Especial, Frágil.
// Cada producto tiene el tamaño en centímetros cúbicos. Y además cada tipo de producto
// requiere un adicional al momento de ser enviado:
// Chico: Ningún adicional.
// Mediano: Requiere un %5 más de espacio
// Grande: Requiere un %20 más de espacio
// Frágil: Requiere un %75 más de espacio
// Especial: Sólo puede ser enviado con productos especiales

// Para ello requerimos que los productos guarden el tamaño y tengan un metodo TamanoTotal
// que nos devuelva el el espacio en cm3 que requerimos para ser enviado.

// Y una estructura Flete que tenga los métodos:
// - AgregarProducto: agregar producto al flete
// - CalcularEnvios: calcula la cantidad de envíos que debe realizar sabiendo que solo
// puede cargar un total de 10.000.000 cm3 por envío

type IProducto interface {
	TamanoTotal() float64
}

type Producto4 struct {
	tamano float64
}
type Chico4 struct {
	p Producto4
}
type Mediano4 struct {
	p Producto4
}
type Grande4 struct {
	p Producto4
}
type Fragil4 struct {
	p Producto4
}
type Especial4 struct {
	p Producto4
}

func (c Chico4) TamanoTotal() float64 {
	return c.p.tamano
}
func (c Mediano4) TamanoTotal() float64 {
	return c.p.tamano * 1.05
}
func (c Grande4) TamanoTotal() float64 {
	return c.p.tamano * 1.20
}
func (c Fragil4) TamanoTotal() float64 {
	return c.p.tamano * 1.75
}
func (c Especial4) TamanoTotal() float64 {
	return c.p.tamano
}

// Y una estructura Flete que tenga los métodos:
// - AgregarProducto: agregar producto al flete
// - CalcularEnvios: calcula la cantidad de envíos que debe realizar sabiendo que solo
// puede cargar un total de 10.000.000 cm3 por envío
type Flete struct {
	productos []IProducto
}

func (f *Flete) AgregarProducto(producto IProducto) {
	f.productos = append(f.productos, producto)
}

func (f Flete) CalcularEnvios() (int, int) {
	var cantidad float64 = 10000000
	var accum, accumEsp float64
	var envios int = 1
	var enviosEsp int

	for _, producto := range f.productos {
		if reflect.TypeOf(producto).String() == "main.Especial4" {
			accumEsp += producto.TamanoTotal()
		}
		accum += producto.TamanoTotal()
	}
	for accum > cantidad {
		accum -= cantidad
		envios += 1
	}

	if accumEsp != 0 {
		enviosEsp = 1
	}
	for accumEsp > cantidad {
		accumEsp -= cantidad
		enviosEsp += 1
	}

	return envios, enviosEsp
}
