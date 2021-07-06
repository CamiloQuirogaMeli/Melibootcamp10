package main

import (
	"fmt"
)

func impuesto(sueldo float64) (imp float64) {

	switch {
	case sueldo > 50000 && sueldo <150001:
		imp = 0.17
	case sueldo > 150000:
		imp = 0.27
	default:
		imp = 0
	}

	return
}

func main(){
	var sueldo float64

	sueldo = 50000

	imp := impuesto(sueldo)

	fmt.Println("El sueldo del empleado es de: $", sueldo, ".\nSe le aplicará un descuento del ", imp*100, "%.")
	fmt.Println("Descuento total aplicado: ", sueldo*imp)
	fmt.Println("Total a pagar: ", sueldo-(sueldo*imp))
}




/* Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de
depositar el sueldo, para cumplir el objetivo necesitan una función que nos devuelva el
impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un %17 del
sueldo y si gana más de $150.000 se le descontará además un %10. */