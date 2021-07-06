package main

import "fmt"

type Ecommerce interface {
	Precio() float64
	Envio() string
}

type TiendaUno struct {
	Ubicacion string
	Tipo      string
	Costo     float64
	Direccion string
}

type TiendaDos struct {
	CantProductos float64
	Tipo          string
	Costo         float64
	Direccion     string
}

func nuevaTienda(p string) Ecommerce {
	if p == "TiendaUno" {
		return TiendaUno{Ubicacion: "Lavalle", Tipo: "chico", Costo: 10, Direccion: "BsAs"}
	}
	if p == "TiendaDos" {
		return TiendaDos{CantProductos: 10, Tipo: "chico", Costo: 10, Direccion: "BsAs"}
	}
	return nil
}

func (p TiendaDos) Precio() float64 {
	if p.Tipo == "chico" {
		return p.Costo
	} else if p.Tipo == "mediano" {
		return p.Costo * 1.03
	} else {
		return (p.Costo * 1.06) + 2500
	}
}

func (p TiendaDos) Envio() string {
	return p.Direccion
}

func (p TiendaUno) Precio() float64 {
	if p.Tipo == "chico" {
		return p.Costo
	} else if p.Tipo == "mediano" {
		return p.Costo * 1.03
	} else {
		return (p.Costo * 1.06) + 2500
	}
}

func (p TiendaUno) Envio() string {
	return p.Direccion
}

func main() {
	p := nuevaTienda("TiendaDos")
	var pre = p.Precio()
	var env = p.Envio()

	fmt.Println("el precio final es", pre)
	fmt.Println("el envio se hace a ", env)
}
