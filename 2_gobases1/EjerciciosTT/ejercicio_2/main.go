package main

import (
	"fmt"
)

func main() {

	fmt.Println("Enter the Price of the Product: ")
	var price float64
	fmt.Scanln(&price)

	fmt.Println("Enter the Discount for the Product: ")
	var discount float64
	fmt.Scanln(&discount)

	fmt.Println("Price original: ", price)
	fmt.Println("Discount: ", discount)
	fmt.Println("Precio final: ", (price - (price * discount / 100)))

}
