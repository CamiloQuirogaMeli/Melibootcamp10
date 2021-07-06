package main

import "fmt"

func priceWithDiscount(price float64, discount int) float64 {
	discountPrice := price * float64(discount) / 100.0
	return price - discountPrice
}

func main() {
	var price float64
	var discount int

	price = 120.0
	discount = 30

	fmt.Println("Precio inicial: ", price)
	fmt.Println("Descuento aplicado: ", discount)
	fmt.Println("Precio final: ", priceWithDiscount(price, discount))
}