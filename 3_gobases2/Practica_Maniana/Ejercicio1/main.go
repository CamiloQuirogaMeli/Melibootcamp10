package main

import "fmt"

func calculoImpuesto(importe float64) float64 {
	var impuesto float64

	if importe > 150000 {
		impuesto = importe * 0.27
	} else if importe > 50000 {
		impuesto = importe * 0.17
	} else {
		impuesto = 0
	}
	return impuesto
}

func main() {
	fmt.Println("El impuesto es: ", calculoImpuesto(160000))
}
