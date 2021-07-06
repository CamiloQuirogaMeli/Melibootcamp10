package main

import f "fmt"

func main() {
	/*
		Ejercicio 1 - Impuestos de salario

		Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de
		depositar el sueldo, para cumplir el objetivo necesitan una función que nos devuelva el
		impuesto de un salario.
		Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un %17 del
		sueldo y si gana más de $150.000 se le descontará además un %10.
	*/

	var salary float64
	f.Println("Ingrese salario")
	f.Scanln(&salary)
	imp := desc(salary)
	f.Println("El descuento a realizar es de $", imp, "\nEn sueldo liquido es de $", salary-imp)
}

func desc(salary float64) float64 {
	var discount float64 = 0
	if salary > 50000 {
		discount = (salary * 17) / 100
	}
	if salary > 150000 {
		discount += (salary * 10) / 100
	}
	return discount
}
