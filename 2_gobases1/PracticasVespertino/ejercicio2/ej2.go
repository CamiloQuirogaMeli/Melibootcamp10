package main

import "fmt"

var (
	price    = 100
	discount = 30
)

func main() {

	fmt.Println("El producto tiene el precio de: ", price)
	fmt.Println("El descuento es de: ", discount)
	fmt.Println("El valor con el dto. aplicado es de: ", ((price * (100 - discount)) / 100))
}

/*
	 Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos, para ello
	necesitan una aplicación que les permita calcular el descuento en base a 2 variables, su
	precio y el descuento en porcentaje. Espera obtener como resultado el valor con el
	descuento aplicado y luego imprimirlo en consola.
*/
