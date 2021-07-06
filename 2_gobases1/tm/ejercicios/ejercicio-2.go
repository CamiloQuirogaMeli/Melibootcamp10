package ejercicios

import "fmt"

func SegundoEjercicio() {
	var (
		price = 1000
		off   = 10
	)

	fmt.Printf("El precio del articulo es: %d\n", price)
	fmt.Printf("El descuento aplicado es: %d%%\n", off)
	fmt.Printf("El precio con descuento es: %d\n", price-(price/100)*off)
}
