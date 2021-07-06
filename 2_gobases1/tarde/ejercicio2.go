package main

import "fmt"

func main() {
	price := 2000
	discount_percentage := 0.25

	discount := float64(price) * discount_percentage

	fmt.Println("Este producto cuesta $", price, "pero tiene un", discount_percentage*100, " % de descuento!!")
	fmt.Println("Por lo tanto el descuento del producto seria de $", discount)
}
