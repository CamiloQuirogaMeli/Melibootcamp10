package main

import "fmt"

func main() {
	impuestoEmpleado1 := impuestoDeSalario(25000)
	impuestoEmpleado2 := impuestoDeSalario(70000)
	impuestoEmpleado3 := impuestoDeSalario(160000)
	fmt.Println(impuestoEmpleado1)
	fmt.Println(impuestoEmpleado2)
	fmt.Println(impuestoEmpleado3)
}

func impuestoDeSalario(salario float64) float64 {
	switch {
	case salario <= 50000:
		return 0
	case salario <= 150000:
		return salario * 0.17
	default:
		return salario * 0.27
	}
}
