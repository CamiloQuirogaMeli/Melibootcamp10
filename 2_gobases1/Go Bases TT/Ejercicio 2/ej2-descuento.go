package main

import "fmt"

func main() {
	var precio, descuento float32

	fmt.Print("Ingrese el precio: ")
	fmt.Scanln(&precio)

	fmt.Print("Porentaje de descuento: ")
	fmt.Scanln(&descuento)

	fmt.Println("Precio final: ", precio-precio*descuento/100)
}
