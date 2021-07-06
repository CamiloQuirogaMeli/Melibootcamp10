package main

import (
	"fmt"
)

func main() {
	/*
		Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de
		depositar el sueldo, para cumplir el objetivo necesitan una función que nos devuelva el
		impuesto de un salario.
		Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un %17 del
		sueldo y si gana más de $150.000 se le descontará además un %10.
	*/
	var salario float64

	fmt.Println("Ingresa el salario: ")
	fmt.Scanln(&salario)
	descuento, porcentajeDescuento := CalcularImpuesto(salario)
	fmt.Println("El porcentaje de descuento es: ", porcentajeDescuento)
	fmt.Println("El descuento total es:", descuento)
	fmt.Println("El sueldo neto es:", salario-descuento)
}

func CalcularImpuesto(salario float64) (float64, float64) {

	var descuento float64
	var porcentajeDescuento float64

	if salario > 0 {
		if salario < 50000 {
			porcentajeDescuento = 0.17
		} else {
			porcentajeDescuento = 0.27
		}
	} else {
		fmt.Println("El salario no puede ser negativo")
		porcentajeDescuento = 0
		panic("Número negativo")
	}
	descuento = salario * porcentajeDescuento
	return descuento, porcentajeDescuento
}
