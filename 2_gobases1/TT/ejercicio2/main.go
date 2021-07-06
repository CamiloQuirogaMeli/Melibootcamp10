package main

import "fmt"

func main() {
	descuento(300.0, 0.10)
}

func descuento(precio float64, descuento float64) {

	var preciofinal float64 = precio - (precio * descuento)

	fmt.Println(preciofinal)

}
