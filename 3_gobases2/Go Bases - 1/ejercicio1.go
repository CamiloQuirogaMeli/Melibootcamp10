package main

import "fmt"

/*
	Ejercicio 1 - Impuestos de salario
	Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo,
	para cumplir el objetivo necesitan una función que nos devuelva el impuesto de un salario.
	Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un %17 del sueldo y si gana más
	de $150.000 se le descontará además un %10.
*/

func impuestoSalario(salario float64) float64 {
	if salario >= 50000 && salario <= 150000 {
		salario = salario - (salario * 0.17)
	} else if salario > 150000 {
		salario = salario - (salario * 0.27)
	} else {
		salario = salario - 0
	}
	return salario
}

func main() {
	salario := 0.0
	fmt.Print("Digite el salario del empleado: ")
	fmt.Scan(&salario)
	fmt.Println("El salario del empleado con impuesto es:", impuestoSalario(salario))
}
