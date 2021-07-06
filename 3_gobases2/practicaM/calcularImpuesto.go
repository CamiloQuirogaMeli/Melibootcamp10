package main

import (
	"fmt"
)

func impuestoSalario(salario float64) float64 {
	var impuesto, descuento float64
	if salario >= 50000 {
		descuento = 0.17
		impuesto = salario * descuento

		if salario >= 150000 {
			descuento = 0.1
			impuesto += salario * descuento
			return impuesto
		}
		return impuesto
	}
	return 0
}

func main() {
	fmt.Println("El impuesto a descontar es de: $", impuestoSalario(150000))
}
