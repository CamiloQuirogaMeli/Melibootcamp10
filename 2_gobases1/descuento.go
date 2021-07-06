package main

import "fmt"

//Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos, para ello necesitan una aplicación que les permita calcular el descuento en base a 2
//variables, su precio y el descuento en porcentaje. Espera obtener como resultado el valor con el descuento aplicado y luego imprimirlo en consola.
//● Crear la aplicación de acuerdo a los requerimientos.

func main()  {
	var price, desc = 25990, 25

	fmt.Println("El producto con valor de $", price, ", tiene un descuento vigente del ", desc, "%, y el valor total sería el siguiente:")
	fmt.Println("$", (price * (100 - desc)) / 100)
}