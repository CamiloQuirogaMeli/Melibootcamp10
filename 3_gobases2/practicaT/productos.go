package main

import "fmt"

type eCommerce interface {
	precio() float64
	envio() string
}

type tiendaUno struct {
	direccion, nombre, tipoProducto string
}

type tiendaDos struct {
	direccion, nombre, tipoProducto string
}

func (tUno tiendaUno) precio() float64 {
	switch tUno.tipoProducto {
	case "pequeño":
		return 5000
	case "mediano":
		var costoMediano float64 = 10000
		interes := costoMediano * 0.03
		costoMediano += interes
		return costoMediano
	case "grande":
		var costoGrande float64 = 15000
		interes := costoGrande * 0.06
		costoGrande += interes + 2500
		return costoGrande
	}
	return 0
}

func (tUno tiendaUno) envio() string {
	return tUno.direccion
}

func (tDos tiendaDos) precio() float64 {
	switch tDos.tipoProducto {
	case "pequeño":
		return 300
	case "mediano":
		var costoMediano float64 = 1000
		interes := costoMediano * 0.03
		costoMediano += interes
		return costoMediano
	case "grande":
		var costoGrande float64 = 3000
		interes := costoGrande * 0.06
		costoGrande += interes + 2500
		return costoGrande
	}
	return 0
}

func (tDos tiendaDos) envio() string {
	return tDos.direccion
}

func nuevaTienda(tienda string, tipoProducto string) eCommerce {
	switch tienda {
	case "tiendaUno":
		return tiendaUno{direccion: "Calle falsa 123", nombre: "Tienda Uno", tipoProducto: tipoProducto}
	case "tiendaDos":
		return tiendaDos{direccion: "Springfield", nombre: "Tienda Dos", tipoProducto: tipoProducto}
	}
	return nil
}

func main() {
	const (
		SMALL  = "pequeño"
		MEDIUM = "mediano"
		BIG    = "grande"
	)

	t1 := nuevaTienda("tiendaUno", SMALL)
	fmt.Println("Tienda uno")
	fmt.Println(t1.precio())
	fmt.Println(t1.envio())
	t2 := nuevaTienda("tiendaDos", SMALL)
	fmt.Println("Tienda dos")
	fmt.Println(t2.precio())
	fmt.Println(t2.envio())
}
