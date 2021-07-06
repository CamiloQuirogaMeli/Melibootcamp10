package main

import "fmt"

func incomeTax(income float64) float64 {
	var tax float64
	switch{
		case income > 150000:
			tax = (0.17 + 0.10) * income
		case income > 50000:
			tax = 0.17 * income
		default:
			tax = 0.0
	}
	return tax
}

func main() {

	for i := 0; i < 3; i++ {

		var income float64 = 10000.0 + float64(i) * 100000.0
		fmt.Printf("Salario: %.2f\tImpuesto: %.2f\n", income, incomeTax(income))
	}
}