package main

import (
	f "fmt"
)

const (
	CHICO    = "chico"
	MEDIANO  = "mediano"
	GRANDE   = "grande"
	TIENDA_1 = "tiendaUno"
	TIENDA_2 = "tiendaDos"
)

type tiendaUno struct {
	Producto producto
}

type tiendaDos struct {
	Producto producto
}

type producto struct {
	Nombre         string
	DireccionEnvio string
	Precio         float32
	Tipo           string
}

type ecommerce interface {
	Precio() string
	Envio() string
}

func (tienda tiendaUno) Precio() string {
	switch tienda.Producto.Tipo {
	case CHICO:
		return "El producto: " + tienda.Producto.Nombre + " tiene un precio de: " + f.Sprint(tienda.Producto.Precio)
	case MEDIANO:
		return "El producto: " + tienda.Producto.Nombre + " tiene un precio de: " + f.Sprint(tienda.Producto.Precio*1.03) + " incluido el costo de mantenimiento"
	case GRANDE:
		return "El producto: " + tienda.Producto.Nombre + " tiene un precio de: " + f.Sprint(tienda.Producto.Precio*1.06+2500.0) + " incluido el costo de mantenimiento y cuota de envio por tamano"
	default:
		return "El producto: " + tienda.Producto.Nombre + " tiene un precio de: " + f.Sprint(tienda.Producto.Precio)
	}
}

func (tienda tiendaUno) Envio() string {
	return "El producto:" + tienda.Producto.Nombre + " se enviara a la dirección: " + tienda.Producto.DireccionEnvio
}

func (tienda tiendaDos) Precio() string {
	switch tienda.Producto.Tipo {
	case CHICO:
		return "El producto: " + tienda.Producto.Nombre + " tiene un precio de: " + f.Sprint(tienda.Producto.Precio)
	case MEDIANO:
		return "El producto: " + tienda.Producto.Nombre + " tiene un precio de: " + f.Sprint(tienda.Producto.Precio*1.03) + " incluido el costo de mantenimiento"
	case GRANDE:
		return "El producto: " + tienda.Producto.Nombre + " tiene un precio de: " + f.Sprint(tienda.Producto.Precio*1.06+2500.0) + " incluido el costo de mantenimiento y cuota de envio por tamano"
	default:
		return "El producto: " + tienda.Producto.Nombre + " tiene un precio de: " + f.Sprint(tienda.Producto.Precio)
	}
}

func (tienda tiendaDos) Envio() string {
	return "El producto:" + tienda.Producto.Nombre + " se enviara a la dirección: " + tienda.Producto.DireccionEnvio
}

func nuevaTienda(tienda string, tipoProducto string) ecommerce {
	switch tienda {
	case TIENDA_1:
		return tiendaUno{Producto: producto{Nombre: "Producto 1", DireccionEnvio: "Direccion 1", Precio: 100, Tipo: tipoProducto}}
	case TIENDA_2:
		return tiendaDos{Producto: producto{Nombre: "Producto 2", DireccionEnvio: "Direccion 2", Precio: 200, Tipo: tipoProducto}}
	default:
		return nil
	}
}

func main() {
	producto1 := nuevaTienda(TIENDA_1, MEDIANO)
	f.Println(producto1.Envio())
	f.Println(producto1.Precio())
	producto2 := nuevaTienda(TIENDA_2, GRANDE)
	f.Println(producto2.Envio())
	f.Println(producto2.Precio())
}
