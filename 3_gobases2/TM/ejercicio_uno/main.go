package main

import (
	"fmt"
)

const TOPE_UNO float32 = 50000
const TOPE_DOS float32 = 150000

func calcularImpuesto(salario float32) float32 {

	var impuesto float32
	impuesto = 0

	if salario > TOPE_UNO {
		impuesto = salario * 0.17
	} else if salario > TOPE_DOS {
		impuesto = salario * 0.27
	}

	return impuesto
}

func main() {

	var salario float32
	fmt.Print("Ingrese el salario: ")
	fmt.Scanf("%f", &salario)
	fmt.Println("Salario: ", salario)
	fmt.Print("Impuesto: ", calcularImpuesto(salario))
}
