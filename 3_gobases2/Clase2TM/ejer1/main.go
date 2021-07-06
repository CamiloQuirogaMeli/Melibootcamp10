package main

import "fmt"

func calculateTaxes(salary float64) float64 {

	if salary >= 50000 {
		return salary - (salary * 17 / 100)
	}

	if salary >= 150000 {
		return salary - (salary * 17 / 100) - (salary * 10 / 100)
	}

	return salary
}

func main() {

	var salary float64 = 160000.0
	totalSalary := calculateTaxes(salary)

	fmt.Println(totalSalary)
}
