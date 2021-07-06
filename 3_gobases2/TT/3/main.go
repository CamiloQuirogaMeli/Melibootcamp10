package main

import (
	"errors"
	"fmt"
)

const (
	CHICO   = "chico"
	GRANDE  = "grande"
	MEDIANO = "mediano"
)

func main() {

	tienda, err := nuevaTienda("grande")

	if err == nil {
		precio := tienda.Precio(10)
		envio := tienda.Envio()

		fmt.Println("Direccion :", envio)
		fmt.Println("Precio producto: ", precio)
	} else {
		fmt.Println(err)
	}
}

type Ecommerce interface {
	Precio(float64) float64
	Envio() string
}

type tiendaMayorista struct {
	nombre    string
	direccion string
	sitio     string
}

type tiendaMinorista struct {
	nombre    string
	direccion string
	sitio     string
	producto  string
}

func nuevaTienda(t string) (Ecommerce, error) {

	switch t {
	case CHICO, MEDIANO:
		return tiendaMinorista{nombre: "MinoristaShop", direccion: "calle rara 123", sitio: "Minorista.com", producto: t}, nil
	case GRANDE:
		return tiendaMayorista{nombre: "MayoristaShop", direccion: "calle mas rara 345", sitio: "MayoristaShop.com"}, nil
	default:
		return nil, errors.New("Tipo de producto no registrado")
	}
}

func (t tiendaMinorista) Precio(p float64) float64 {
	if t.producto == MEDIANO {
		return (p + (p * 0.3))
	} else {
		return p
	}
}

func (t tiendaMinorista) Envio() string {
	return t.direccion
}

func (t tiendaMayorista) Precio(p float64) float64 {
	return p + (p * 0.6)
}

func (t tiendaMayorista) Envio() string {
	return t.direccion
}
