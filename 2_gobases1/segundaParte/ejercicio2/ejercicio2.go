package main

import "fmt"

func main() {

	var productPrice, productDiscount, finalPrice int
	fmt.Println("Escribe el precio de tu producto: ")
	fmt.Scanln(&productPrice)
	fmt.Println("Escribe el valor de porcentaje sin % de 0 a 100 en descuento del producto: ")
	fmt.Scanln(&productDiscount)
	finalPrice = productPrice * (100 - productDiscount) / 100
	fmt.Printf("El valor final con descuento es %d\n", finalPrice)

}
