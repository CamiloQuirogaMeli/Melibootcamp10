package main

import "fmt"

// Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de
// depositar el sueldo, para cumplir el objetivo necesitan una función que nos devuelva el
// impuesto de un salario.
// Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un %17 del
// sueldo y si gana más de $150.000 se le descontará además un %10.
func calcTax() float64 {
	var salary, tax float64
	fmt.Println("Ingrese su salario: ")
	fmt.Scan(&salary)

	if salary > 50000 {
		tax += 0.17
	}
	if salary > 150000 {
		tax += 0.1
	}

	return salary * (1 - tax)
}
