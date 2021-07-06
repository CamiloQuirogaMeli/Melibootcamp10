package main

import "fmt"

func main() {
	fmt.Printf("el impuesto del salario es %.2f\n", taxes(100000))
}

func taxes(salary float64) float64 {
	switch {
	case salary > 150000:
		return salary * 0.1
	case salary > 50000:
		return salary * 0.17
	default:
		return 0
	}
}
