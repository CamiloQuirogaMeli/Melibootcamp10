package main

import "fmt"

func salaryTaxes(salary float64) float64 {
	var taxesTotal float64 = 0.000
	if salary > 50000 {
		taxesTotal = 0.17
		if salary > 150000 {
			taxesTotal = taxesTotal * 1.10
		}

	}

	return taxesTotal

}

func main() {
	currentSalary := 0.0
	fmt.Println("Ingrese su sueldo para calcular su impuesto")
	fmt.Scanln(&currentSalary)
	finalTaxes := salaryTaxes(currentSalary)
	fmt.Println("Si su sueldo es de $", currentSalary)
	fmt.Printf("Su impuesto sera de %.2f %% \n", finalTaxes)
	fmt.Printf("Y su sueldo neto sera de %.2f $ \n", currentSalary-currentSalary*finalTaxes)

}
