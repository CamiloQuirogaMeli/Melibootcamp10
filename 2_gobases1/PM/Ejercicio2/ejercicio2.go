package main

import "fmt"

func main() {

	fmt.Println("Ingresá el valor del producto: ")

	var precio float64

	fmt.Scanln(&precio)
	fmt.Println("Ingresá el porcentaje de descuento: ")
	var desc float64
	fmt.Scanln(&desc)

	fmt.Println("Precio: ", precio)
	fmt.Println("Descuento: ", desc)
	fmt.Println("Precio final: ", calcularPrecioFinal(precio, desc))

}

func calcularPrecioFinal(precio float64, descuento float64) float64 {
	return precio - (precio * descuento / 100)
}
