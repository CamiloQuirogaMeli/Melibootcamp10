package main

import "fmt"

func ejercicio8() {

	tienda := nuevaTienda("grande", "murguiondo 1675", 145)
	fmt.Println(tienda)
	fmt.Println(tienda.precio())
	fmt.Println(tienda.envio())
}

const (
	PEQUENO          = "pequeño"
	MEDIANO          = "mediano"
	GRANDE           = "grande"
	ENVIO            = 2500
	IMPUESTO_MEDIANO = 1.003
	IMPUESTO_GRANDE  = 1.006
)

type Tienda struct {
	producto  string
	tipo      string
	valor     float64
	direccion string
}

type TiendaUno struct {
	Tienda
}
type TiendaDos struct {
	Tienda
}

type eCommerce interface {
	precio() float64
	envio() string
}

func (t Tienda) precio() float64 {
	switch t.tipo {
	case PEQUENO:
		return t.valor
	case MEDIANO:
		return t.valor * IMPUESTO_MEDIANO
	case GRANDE:
		return t.valor*IMPUESTO_GRANDE + ENVIO
	}
	return 0
}

func (t Tienda) envio() string {
	return t.direccion
}

func nuevaTienda(tipoProducto string, direccion string, valor float64) eCommerce {

	return Tienda{producto: "fideos", tipo: tipoProducto, valor: valor, direccion: direccion}
}
