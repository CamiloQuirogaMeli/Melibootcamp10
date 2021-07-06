/*
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de
depositar el sueldo, para cumplir el objetivo necesitan una función que nos devuelva el
impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un %17 del
sueldo y si gana más de $150.000 se le descontará además un %10.
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Por favor ingrese su salario neto sin comas(,) ni puntos(.)")
	var salary int
	_, errSalary := fmt.Scanf("%d", &salary)

	if errSalary != nil {
		fmt.Println("Bad Input")
		os.Exit(0)
	}
	discount, salaryDiscount, totalSalary := calcularSalario(salary)

	fmt.Println("Su salario es de: ", salary)
	fmt.Println("Su descuento es de: ", discount, "%")
	fmt.Println("El descuento a su salario es de: ", salaryDiscount)
	fmt.Println("El salario que usted recibiria seria de: ", totalSalary)
}

func calcularSalario(salary int) (discount int, salaryDiscount int, totalSalary int) {

	if salary >= 150000 {
		discount = 27
	} else if salary >= 50000 {
		discount = 17
	} else if salary < 50000 {
		discount = 0
	}

	salaryDiscount = (salary * discount) / 100

	totalSalary = salary - salaryDiscount

	return
}
