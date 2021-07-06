package main

import "fmt"

func main() {
	baseSalary := 160000

	tax := salaryTaxes(baseSalary)
	fmt.Println("El impuesto al salario", baseSalary, "es", tax)
}
