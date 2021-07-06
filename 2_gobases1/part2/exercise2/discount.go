package main

import (
	"fmt"
)

func main() {
	var price float32
	var discount float32
	var total int

	fmt.Print("Enter item price: ")
	fmt.Scanln(&price)
	fmt.Print("Enter the discount % : ")
	fmt.Scanln(&discount)
	total = int(price * (1 - (discount / 100)))
	fmt.Printf("Original item price: $%.2f\n", price)
	fmt.Printf("Item price with a discount of %.0f percent: $%d\n", discount, total)
}