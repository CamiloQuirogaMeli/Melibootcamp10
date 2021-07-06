package main

import "fmt"

func main() {
	var price int = 1500
	var disc int = 30
	fmt.Printf("Precio: %d\n", price)
	fmt.Printf("Descuento: %d\n", disc)
	fmt.Printf("Precio con descuento: %d\n", (100-disc)*price/100)
}
