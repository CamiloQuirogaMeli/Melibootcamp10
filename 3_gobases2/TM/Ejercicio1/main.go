package main

import (
	"fmt"
	"math"
)

func main() {
	Ejercicio1()
}

func Ejercicio1() {
	fmt.Println("Ejercicio1")
	var (
		salario           = 75000.00
		impuestoBase      = 17.00
		impuestoAdicional = 10.00
	)
	if salario > 150000.00 {
		fmt.Println("El impuesto que se debe cobrar es:", impuestoBase+impuestoAdicional, "%. El valor del impuesto es: $", math.Round(salario*((impuestoBase+impuestoAdicional)/100)))
	} else if salario <= 0.00 {
		fmt.Println("El salario no puede ser $0.00 o menor")
	} else {
		fmt.Println("El impuesto que se debe cobrar es:", impuestoBase, "%. El valor del impuesto es: $", math.Round(salario*(impuestoBase/100)))
	}
}
