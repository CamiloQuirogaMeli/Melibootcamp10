package main

import (
	"fmt"
)

func main() {

	discount_f := 30.5

	price_f := 100.50

	var discount_total float64 = price_f * (discount_f / 100)
	var price_total float64 = price_f - discount_total

	fmt.Printf("El precio es: $ %v \n", price_f)
	fmt.Printf("El  descuento es: $ %v \n", discount_total)
	fmt.Printf("El precio a abonar es: $ %v \n", price_total)
}
