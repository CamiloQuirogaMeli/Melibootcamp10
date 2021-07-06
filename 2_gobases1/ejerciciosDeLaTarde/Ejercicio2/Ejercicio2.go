package main

import "fmt"

func main() {
	var discount float64
	var price float64

	fmt.Printf("Ingrese el precio ")
	fmt.Scanf("%f", &price)

	fmt.Printf("Ingrese el descuento ")
	fmt.Scanf("%f", &discount)

	fmt.Printf("El precio con descuento es de %f ", price-(discount*price)/100)

}
