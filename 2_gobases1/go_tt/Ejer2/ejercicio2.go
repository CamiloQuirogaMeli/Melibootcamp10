/*
Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos, para ello
necesitan una aplicación que les permita calcular el descuento en base a 2 variables, su
precio y el descuento en porcentaje. Espera obtener como resultado el valor con el
descuento aplicado y luego imprimirlo en consola.
*/

// package declaration
package main

// library imports
import (
	"fmt"
	"os"
)

func main() {
	welcome := `----------------------------------------------------------------
----------------- Bienvenidos a su tienda virtual --------------
----------------------------------------------------------------`
	fmt.Println(welcome)
	fmt.Print("por favor digite el precio del producto sin puntos (.)  ni comas (,) : ")
	//Create variable for read console
	var product, discount int
	_, errProduct := fmt.Scanf("%d", &product)

	if errProduct != nil {
		fmt.Println("Bad Input")
		os.Exit(0)
	}

	fmt.Println()

	fmt.Print("por favor digite el descuento del producto sin ( % ): ")

	_, errDiscount := fmt.Scanf("%d", &discount)
	if errDiscount != nil {
		fmt.Println("Bad Input")
		os.Exit(0)
	}

	fmt.Println()
	result := (product * discount) / 100

	fmt.Println("El producto con precio ", product, " USD tiene un descuento del ", discount, "%")
	fmt.Println("El descuento es de ", result, " USD por lo tanto el precio total del producto es ", product-result, " USD")

}
