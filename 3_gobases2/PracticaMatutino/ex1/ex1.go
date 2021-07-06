package main

import "fmt"

func main() {
	salary := 52000.45

	tax := getTaxes(salary)

	fmt.Println("El impuesto es de ", tax)
}

func getTaxes(salary float64) int {

	if salary < 50000.0 {
		return 0
	} else if salary > 150000.0 {
		return 27
	} else {
		return 17
	}
}

/*
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de
depositar el sueldo, para cumplir el objetivo necesitan una función que nos devuelva el
impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un %17 del
sueldo y si gana más de $150.000 se le descontará además un %10.
*/
