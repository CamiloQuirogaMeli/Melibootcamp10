package main

import "fmt"

func main() {
	sueldo := 200000.0

	var impuesto float64 = calcularImpuesto(sueldo)

	fmt.Printf("Se aplica un impuesto de %f", impuesto)
}

func calcularImpuesto(sueldo float64) float64 {
	if sueldo > 150000 {
		return 27
	}
	if sueldo > 50000 {
		return 17
	}
	return 0
}
