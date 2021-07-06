package main

import "fmt"

func calcularSalario(minutosTrabajados int, sector string) (salario float64) {
	switch sector {
	case "A":
		salario = float64(3000 * (minutosTrabajados / 60))
		salario = salario + (salario / 2)
	case "B":
		salario = float64(1500 * (minutosTrabajados / 60))
		salario = salario + (salario * 0.2)
	case "C":
		salario = float64(1000 * (minutosTrabajados / 60))
	}
	return salario
}

func main() {
	fmt.Printf("El salario del empleado es de: %.2f\n", calcularSalario(60, "A"))
}
