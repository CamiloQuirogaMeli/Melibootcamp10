package main

import (
	"fmt"
)

func main() {
	var descuento = 20 //Eliga el descuento
	var precio = 100   //Eliga el precio
	var precioFinal = precio - precio*descuento/100
	fmt.Println("El precio del producto es de: $", precio)
	fmt.Println("El descuento elegido es de: ", descuento)
	fmt.Println("El precio final del producto con el descuento aplicado es: $", precioFinal)

}
