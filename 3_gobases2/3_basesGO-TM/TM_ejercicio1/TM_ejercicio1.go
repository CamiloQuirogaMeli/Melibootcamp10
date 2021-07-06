package main

import "fmt"

/*
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de
depositar el sueldo, para cumplir el objetivo necesitan una función que nos devuelva el
impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un %17 del
sueldo y si gana más de $150.000 se le descontará además un %10.
*/

/* Calculo el impuesto en función del salario del empleado */
func calculaImpuesto(salario float64) float64 {
	var impuesto float64

	switch {
	case salario > 150000.0:
		impuesto = 0.27
	case salario > 50000.0:
		impuesto = 0.17
	default:
		impuesto = 0

	}

	return impuesto

}

func main() {

	var salario float64 = 100000

	imp := calculaImpuesto(salario)

	fmt.Println("El salario bruto del empleado es: ", salario)
	fmt.Println("El impuesto aplicado al salario es de: ", imp*100, "%")
	fmt.Println("Sueldo neto: ", salario-salario*imp)

}
