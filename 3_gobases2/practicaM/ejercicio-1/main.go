package main

import "fmt"

func main() {
	fmt.Println("Se le descontará al empleado $", taxSalary(150001))
}

func taxSalary(salary float64) float64 {
	switch {
	case salary > 50000:
		return salary * 0.17
	case salary > 150000:
		return salary * 0.27
	default:
		return 0
	}
}
