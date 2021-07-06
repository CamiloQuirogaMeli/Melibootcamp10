package main

import "fmt"

func main() {
	salary := 160000.0
	fmt.Println("Salary without taxes is ", salary)
	salary = getSalaryWithTaxes(salary)
	fmt.Println("Salary with taxes is ", salary)
}

func getSalaryWithTaxes(salary float64) float64 {

	if salary > 150000 {
		return salary * (1 - 0.27)
	} else if salary > 50000 {
		return salary * (1 - 0.17)
	} else {
		return salary
	}
}
