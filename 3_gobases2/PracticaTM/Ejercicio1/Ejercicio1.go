// Ejercicio 1 - Impuestos de salario

// Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de
// depositar el sueldo, para cumplir el objetivo necesitan una función que nos devuelva el
// impuesto de un salario.
// Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un %17 del
// sueldo y si gana más de $150.000 se le descontará además un %10.

package main

import "fmt"

func calculateTax(salary float64) float64 {
	var tax float64 = 0

	if salary > 50000 && salary <= 150000 {
		tax = salary * 0.17
	} else if salary > 150000 {
		tax = salary * 0.27
	}

	return tax
}

func main() {

	var salary float64
	var tax float64

	fmt.Printf("Ingrese el sueldo de su empleado: ")
	fmt.Scan(&salary)
	tax = calculateTax(salary)
	fmt.Printf("El impuesto para el empleado es de: %.2f", tax)

}
