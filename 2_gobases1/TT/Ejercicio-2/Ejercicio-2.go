package main

import "fmt"

func main() {
	fmt.Println("Ejercicio 2")
	var originalPrice, discountPercentage float32
	fmt.Println("Ingrese el monto del articulo: ")
	fmt.Scanln(&originalPrice)
	fmt.Println("Ingrese el porcentaje del descuento")
	fmt.Scanln(&discountPercentage)
	fmt.Println("El precio final del producto es: ", originalPrice*discountPercentage/100)
}
