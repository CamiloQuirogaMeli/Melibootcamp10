package main

import "fmt"

func main() {
	var salary float64 = 160000
	fmt.Printf("El impuesto por un salario de $%.2f es de: $%.2f\n", salary, salaryTax(salary))
}

func salaryTax(salary float64) float64 {
	var tax float64

	if salary > 50000 {
		tax += salary * .17
	}

	if salary > 150000 {
		tax += salary * .1
	}

	return tax
}
