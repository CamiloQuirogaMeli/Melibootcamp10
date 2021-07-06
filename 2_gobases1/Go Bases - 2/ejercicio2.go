package main

import "fmt"

/*
	Ejercicio 2 - Descuento
		Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos,
		para ello necesitan una aplicación que les permita calcular el descuento en base a
		2 variables, su precio y el descuento en porcentaje. Espera obtener como resultado
		el valor con el descuento aplicado y luego imprimirlo en consola.
*/

func main() {
	var precio, descuento float64
	fmt.Print("Introduzca el precio del producto: ")
	fmt.Scan(&precio)
	fmt.Print("Introduzca el descuento del producto: ")
	fmt.Scan(&descuento)
	precio = precio - (precio * (descuento / 100))
	fmt.Printf("El valor del producto con el descuento aplicado es: %g\n", precio)
}
