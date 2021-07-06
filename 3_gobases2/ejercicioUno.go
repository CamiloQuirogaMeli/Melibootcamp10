package main

import "fmt"

func calcularImpuesto(sueldo float64) float64 {
	var descuento float64 = 0
	if sueldo > 50000 {
		descuento = 0.17 * sueldo
	} else if sueldo > 150000 {
		descuento += 0.1 * sueldo
	}
	return float64(descuento)
}

func main() {
	var sueldo float64 = 440000

	fmt.Printf("El descuento del sueldo sera de $%.2f\n", calcularImpuesto(sueldo))
}
