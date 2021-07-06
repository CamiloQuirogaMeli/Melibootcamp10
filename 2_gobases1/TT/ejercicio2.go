package main

import "fmt"

func main() {

	var price float32 = 200
	var discount float32 = 25
	var totalPrice float32 = price * (1 - (discount / 100))

	fmt.Printf("El precio con descuento del %.f%% es: $%.2f\n", discount, totalPrice)
}
