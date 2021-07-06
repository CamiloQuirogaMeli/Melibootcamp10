package main

import "fmt"

func main() {

	//
	// Ejercicio 2 - Descuento
	// Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos, para ello
	// necesitan una aplicación que les permita calcular el descuento en base a 2 variables, su
	// precio y el descuento en porcentaje. Espera obtener como resultado el valor con el
	// descuento aplicado y luego imprimirlo en consola.
	// ● Crear la aplicación de acuerdo a los requerimientos.
	//

	precio := 50
	descuento := 200
	precioFinal := precio - ((precio * descuento) / 100)

	fmt.Printf("Precio del producto: $%d\n", precio)
	fmt.Printf("Descuento: %d%%\n", descuento)
	fmt.Printf("Precio final con descuento: $%d\n", precioFinal)

}
