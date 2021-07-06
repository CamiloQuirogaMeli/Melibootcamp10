package main

import (
	"fmt"
)

const (
	CHICO   = "CHICO"
	MEDIANO = "MEDIANO"
	GRANDE  = "GRANDE"
)

type Ecomerce interface {
	Precio(float32) float32
	Envio(string, string) string
}

type Tienda struct {
	nombre      string
	productType string
}

func (tienda Tienda) Precio(costo float32) float32 {

	switch tienda.productType {
	case CHICO:
		return costo
	case MEDIANO:
		return costo + (costo * 0.03)
	case GRANDE:
		return costo + (costo * 0.06) + 2500
	}
	return 0
}

func (tienda Tienda) Envio(direcion, cliente string) string {
	return fmt.Sprint("Estoy enviando un paquete a " + cliente + " con direccion " + direcion + " desde tienda " + tienda.nombre)
}

func nuevaTienda(productType, nombre string) Ecomerce {
	tienda := Tienda{nombre, productType}
	return tienda
}

func main() {

	tienda1 := nuevaTienda(CHICO, "Tienda1")
	tienda2 := nuevaTienda(GRANDE, "Tienda2")

	fmt.Println(tienda1.Envio("Direccion 123", "Cliente1"))
	fmt.Println("Su precio es de ", tienda1.Precio(5000))

	fmt.Println(tienda2.Envio("Direccion 321", "Cliente2"))
	fmt.Println("Su precio es de ", tienda2.Precio(7000))

}
