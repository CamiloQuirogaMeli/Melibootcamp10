package main

import "fmt"

func main() {
	var disc, price float32 = 10, 100
	var totalP float32 = 0
	var totalDisc float32 = 0

	totalDisc = (price * disc) / 100
	totalP = price - totalDisc

	fmt.Println("El valor del producto es de: $", price)
	fmt.Println("Se aplicará un descuento del: ", disc, "%")
	fmt.Println("\nTotal a pagar: $", totalP)
	fmt.Println("Descuento total aplicado: ", totalDisc, "%")

}
