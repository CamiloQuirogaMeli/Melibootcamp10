package main

import "fmt"

func main() {
	/*
	Ejercicio 2 - Descuento
	Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos, para ello
	necesitan una aplicación que les permita calcular el descuento en base a 2 variables, su
	precio y el descuento en porcentaje. Espera obtener como resultado el valor con el
	descuento aplicado y luego imprimirlo en consola.
	● Crear la aplicación de acuerdo a los requerimientos.
	*/

	var precio int
	var descuento int
	fmt.Println("Ingrese precio")
	fmt.Scanln(&precio)
	fmt.Println("Ingrese descuento")
	fmt.Scanln(&descuento)

	fmt.Printf("Precio del producto: $%d\n", precio)
	fmt.Printf("Descuento: %d%%\n", descuento)
	fmt.Printf("Precio final $%d\n", (precio-(precio * descuento)/100))
}
