package main

import "fmt"

func main() {
	var precio, descuento float64 = 100000, 40

	descuento /= 100
	descuento *= precio
	descuento = precio - descuento
	fmt.Println("Valor con descuento aplicado...", descuento)
}
