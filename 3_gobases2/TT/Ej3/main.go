package main

import "fmt"

const (
	small  = "small"
	medium = "medium"
	large  = "large"
)

type Producto struct {
	precio   float64
	category string
}

type tiendaPyM struct {
	nombre   string
	producto Producto
}

type tiendaGrande struct {
	nombre       string
	producto     Producto
	cantClientes int
}

type Ecommerce interface {
	Precio() float64
}

func (t tiendaPyM) Precio() float64 {
	switch t.producto.category {
	case small:
		return t.producto.precio
	case medium:
		return t.producto.precio * 1.03
	default:
		return 0
	}
}

func (t tiendaGrande) Precio() float64 {
	return t.producto.precio*1.06 + 2500
}

func nuevaTienda(tipo string, nombre string, precio float64) Ecommerce {
	switch tipo {
	case small:
		return tiendaPyM{nombre: nombre, producto: Producto{precio, small}}
	case medium:
		return tiendaPyM{nombre: nombre, producto: Producto{precio, medium}}
	case large:
		return tiendaGrande{nombre: nombre, producto: Producto{precio, large}, cantClientes: 0}
	default:
		return nil
	}
}

func main() {
	unaTienda := nuevaTienda(small, "El Sol", 600)
	otraTienda := nuevaTienda(large, "Lalala", 7000)

	fmt.Printf("Tienda 1\n")
	fmt.Printf("Precio (con costo adicional si corresponde): %.2f\n", unaTienda.Precio())
	fmt.Printf("Tienda 2\n")
	fmt.Printf("Precio (con costo adicional si corresponde): %.2f\n", otraTienda.Precio())
}
