package main

import "fmt"

func main() {
	var (
		precio    float64
		descuento float64
	)
	fmt.Println("Ingrese el precio del producto:")
	fmt.Scanln(&precio)
	fmt.Println("Ingrese ahora el descuento a aplicar:")
	fmt.Scanln(&descuento)
	fmt.Println("El precio final es de: ", precio-(precio*descuento/100))
}
