package main

import "fmt"

func main() {
	var precio, descuento float64 = 100.0, 0.4
	fmt.Printf("El precio final con descuento es: %v\n", precio*descuento)
}
