package main

import (
	"fmt"
)

/*
Un Ecommerce necesita realizar una funcionalidad en Go para administrar productos y
retornar el valor del precio total.
La empresa tiene 3 tipos de productos: Chico, Mediano y Grande. (Se espera que sean
muchos más)

Y los costos adicionales son:
- Chico: solo tiene el costo del producto
- Mediano: el precio del producto + un %3 de mantenerlo en la tienda
- Grande: el precio del producto + un %6 de mantenerlo en la tienda y adicional a eso
$2500 de flete.

El porcentaje de mantenerlo en la tienda es en base al precio del producto.

Para ello se requiere una función factory que reciba el tipo de producto y el precio y
retorne una interfaz Producto que tenga el método Precio.

Deber poder ejecutar el método Precio y me traiga el precio total en base al costo del
producto y los adicionales en caso que los tenga
*/

const (
	CHICO   = "chico"
	MEDIANO = "mediano"
	GRANDE  = "grande"
)

type tipo1 struct {
	precioProd float64
}

type tipo2 struct {
	precioProd float64
}

type tipo3 struct {
	precioProd float64
}

type tipoNOT struct {
	mensaje string
}

type producto interface {
	precio() float64
}

func (t tipo1) precio() float64 {
	return t.precioProd
}

func (t tipo2) precio() float64 {
	return t.precioProd * 1.03
}

func (t tipo3) precio() float64 {
	return t.precioProd*1.06 + 2500.0
}

func (t tipoNOT) precio() float64 {
	fmt.Printf(t.mensaje)
	return 1.0
}

func factory(tipoProducto string, precio float64) producto {
	switch tipoProducto {
	case CHICO:
		return tipo1{precioProd: precio}
	case MEDIANO:
		return tipo2{precioProd: precio}
	case GRANDE:
		return tipo3{precioProd: precio}
	default:
		return tipoNOT{mensaje: "tipo de paquete no valido"}

	}
}

func main() {

	prodChico := factory("chico", 5400.0)
	fmt.Println("Precio del producto: ", prodChico.precio())

	prodMediano := factory("mediano", 10000.0)
	fmt.Println("Precio del producto: ", prodMediano.precio())

	prodGrande := factory("grande", 15000.0)
	fmt.Println("Precio del producto: ", prodGrande.precio())

}
