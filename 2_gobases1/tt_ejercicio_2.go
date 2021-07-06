package main

import {
	"fmt"
	"os"
	}

func calcularPrecio(precioActual float32, descuento float32) float32 {
	valor := precioActual * (1 - descuento)
	return valor
}

func main() {
	var precio float32 = os.Args
	var descuento float32 = os.Args

	precioConDescuento := calcularPrecio(precio, descuento)

	fmt.Printf("El precio con el descuento es: %f \n", precioConDescuento)

}
