package main

import (
	"fmt"
)

func main() {

	var (
		precio              int
		descuentoPorcentage int
		precioTotal         int
		descuento           int
	)

	fmt.Print("Ingrese el precio: ")
	fmt.Scanf("%d", &precio)

	fmt.Print("Ingrese el descuento en porcentaje: ")
	fmt.Scanf("%d", &descuentoPorcentage)

	descuento = (precio * descuentoPorcentage) / 100
	precioTotal = precio - descuento

	fmt.Println("el precio total es: ", precioTotal)
}
