// Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos, para ello
// necesitan una aplicación que les permita calcular el descuento en base a 2 variables, su
// precio y el descuento en porcentaje. Espera obtener como resultado el valor con el
// descuento aplicado y luego imprimirlo en consola.
// - Crear la aplicación de acuerdo a los requerimientos.
package main

import "fmt"

func main() {
	var (
		price    float64
		discount float64
	)
	fmt.Printf("Ingrese precio del producto: ")
	fmt.Scanf("%f\n", &price)

	fmt.Printf("Ingrese el porsentaje de descuento: ")
	fmt.Scanf("%f\n", &discount)

	fmt.Println("El precio con el descuento es: ", (1-discount/100)*price)
}
