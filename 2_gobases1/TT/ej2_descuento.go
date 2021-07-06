package main

import (
	"fmt"
)

func desc() {

	fmt.Print("Ingrese su producto: ")
	var product string
	fmt.Scanln(&product)

	fmt.Print("Ingrese el precio del producto ", product, " --> ")
	var productCost float64
	fmt.Scanln(&productCost)

	fmt.Print("Ingrese el descuento que quiere realizar: ")
	var discount float64
	fmt.Scanln(&discount)

	var productWDiscount float64
	productWDiscount = productCost * (1 - (discount / 100))

	fmt.Println("El producto: ", product, " tiene un valor de: ", productCost)
	fmt.Println("Con el descuento aplicado queda en: ", productWDiscount)

}
