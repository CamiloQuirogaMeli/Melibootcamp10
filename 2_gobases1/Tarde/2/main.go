package main

import "fmt"

func main() {
	var price float32 = 100
	var discountPercent float32 = 25
	var discount float32 = (discountPercent * price) / 100
	finalPrice := price - discount

	fmt.Println("Precio final: ", finalPrice)
}
