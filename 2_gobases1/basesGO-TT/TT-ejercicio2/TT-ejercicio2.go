package main

import "fmt"

/*
Ejercicio 2 - Descuento
Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos, para ello
necesitan una aplicación que les permita calcular el descuento en base a 2 variables, su
precio y el descuento en porcentaje. Espera obtener como resultado el valor con el
descuento aplicado y luego imprimirlo en consola.
● Crear la aplicación de acuerdo a los requerimientos.

*/

func valorConDescuento(precio float64, descuento float64) {

	/*Calculo el precio total a abonar en función del descuento aplicado*/

	desc := descuento * precio / 100
	fmt.Println("El precio final a pagar es: ", precio-desc)

}

func main() {

	precio, descuento := 4760.50, 50.0

	fmt.Println("El precio del producto es: ", precio)
	fmt.Println("El descuento aplicado al producto es: ", descuento)

	valorConDescuento(precio, descuento)

}
