package main

import (
	"fmt"
)

func main() {

	/*
		Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos, para ello
		necesitan una aplicación que les permita calcular el descuento en base a 2 variables, su
		precio y el descuento en porcentaje. Espera obtener como resultado el valor con el
		descuento aplicado y luego imprimirlo en consola.
			Crear la aplicación de acuerdo a los requerimientos.
	*/

	var valorInicial float32
	var descuento float32

	fmt.Print("Digita el valor del producto: ")
	fmt.Scanln(&valorInicial)

	fmt.Print("Digita el valor del descuento en porcentajes: ")
	fmt.Scanln(&descuento)

	for descuento > 100 {
		//Se ejecuta siempre que los datos ingresados sean erroneos

		fmt.Println("Descuento no válido")

		fmt.Print("Digita el valor del descuento en porcentajes: ")
		fmt.Scanln(&descuento)

	}

	valorFinal := valorInicial - ((valorInicial * descuento) / 100)

	fmt.Println("El valor del producto con el descuento es de:", valorFinal)

}
