package main

import "fmt"

func main() {
	var precio, descuento, precioFinal float32
	const PORCENTAJE = 100.00

	fmt.Println("Ingrese precio bruto: ")
	fmt.Scanln(&precio)
	fmt.Println("Ingrese % de descuento:")
	fmt.Scanln(&descuento)

	descuento = (precio * descuento) / PORCENTAJE
	precioFinal = precio - descuento

	fmt.Println("El precio final es: ", precioFinal)

}
