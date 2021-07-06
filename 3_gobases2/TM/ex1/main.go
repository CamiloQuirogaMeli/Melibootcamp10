package main

import (
	"fmt"
)

func taxes(salary float64) float64 {
	var percentage float64 = 0
	if salary >= 50000 {
		percentage = 17
	}
	if salary >= 150000 {
		percentage += 10
	}
	return salary / 100.0 * percentage
}

func main() {
	var salary float64
	fmt.Print("Introduzca el salario del empleado: ")
	fmt.Scanln(&salary)
	fmt.Println("El impuesto del salario", salary, "es", taxes(salary))
}
