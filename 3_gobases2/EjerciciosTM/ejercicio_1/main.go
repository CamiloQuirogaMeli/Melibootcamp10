package main

import (
	"fmt"
)

func main() {
	impuesto := calculoImpouesto(56000)

	fmt.Println("El impuesto es: ", impuesto)
}

func calculoImpouesto(salario float64) float64 {
	if salario > 50000 && salario < 150000 {
		return (salario * 0.17)
	} else if salario > 150000 {
		return (salario * 0.10)
	} else {
		return 0
	}
}
