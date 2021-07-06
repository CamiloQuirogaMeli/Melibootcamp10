package main

import (
	"fmt"
)

func main() {
	var price, discount, finalPrice float64
	fmt.Print("Ingrese el precio: ")
	fmt.Scanln(&price)
	fmt.Print("Ingrese el % de descuento: ")
	fmt.Scanln(&discount)
	finalPrice = price * (1 - (discount / 100))
	fmt.Println("El precio con descuento es:", finalPrice)
}
