package main

import "fmt"

func calculateTax(salary float64) float64 {
	var tax float64
	if salary > 50000 && salary <= 150000 {
		tax = salary * 0.17
	} else {
		tax = salary * 0.27
	}
	return tax
}

func main() {
	var tax float64 = calculateTax(75000)
	fmt.Printf("El impuesto de salario es de %.2f\n", tax)
}
