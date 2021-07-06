// Ejercicio 2 - Descuento
// Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos, para ello
// necesitan una aplicación que les permita calcular el descuento en base a 2 variables, su
// precio y el descuento en porcentaje. Espera obtener como resultado el valor con el
// descuento aplicado y luego imprimirlo en consola.
// Crear la aplicación de acuerdo a los requerimientos.

package main

import "fmt"

func main() {

	var price float64
	var desc float64
	var total float64

	fmt.Print("Indicar el precio de la prenda: ")
	fmt.Scan(&price)
	fmt.Print("Indicar el porcentaje de descuento: ")
	fmt.Scan(&desc)
	total = price - (price * (desc / 100))
	fmt.Printf("El precio de la prenda con descuento es: %.2f", total)

}
