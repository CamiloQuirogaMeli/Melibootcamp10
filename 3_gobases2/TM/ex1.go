package main

import "fmt"

func getEmployeeTax(salary float64) float64 {
	switch {
	case salary > 150000:
		return salary * 0.1
	case salary > 50000:
		return salary * 0.17
	default:
		return 0
	}
}

func main() {
	salary := 200000
	fmt.Println("El impuesto del salario es", getEmployeeTax(float64(salary)))
}
