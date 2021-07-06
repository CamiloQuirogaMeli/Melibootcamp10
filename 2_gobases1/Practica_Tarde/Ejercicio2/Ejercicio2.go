package main

import "fmt"

func main() {
	var precio, porc_desc float32 = 100, 10
	var oferta float32 = precio - (precio * (porc_desc / 100))

	fmt.Printf("El precio con descuento es: %v\n", oferta)

}
