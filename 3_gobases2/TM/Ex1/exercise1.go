package main

import (
	"fmt"
)

func main() {
	var salary float64
	fmt.Println("Enter the salary:")
	fmt.Scanln(&salary)

	fmt.Println("Final Salary:", calcTaxes(salary))

}

func calcTaxes(salary float64) float64 {

	switch {
	case salary > 50000 && salary <= 150000:
		return salary * 0.17
	case salary > 150000:
		return salary * 0.27
	default:
		return salary
	}

}
