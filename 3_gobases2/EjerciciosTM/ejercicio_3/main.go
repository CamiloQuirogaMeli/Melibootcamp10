package main

import (
	"fmt"
)

func main() {

	salario_A := calcularSalario(60, "A")
	salario_B := calcularSalario(120, "B")
	salario_C := calcularSalario(180, "C")

	fmt.Println("Salario categoria A: ", salario_A)
	fmt.Println("Salario categoria B: ", salario_B)
	fmt.Println("Salario categoria C: ", salario_C)
}

func calcularSalario(minutos int, categoria string) float64 {
	horasTrabajadas := float64(minutos / 60)

	switch categoria {
	case "A":
		return 3000*horasTrabajadas + (3000 * horasTrabajadas * 0.5)
	case "B":
		return 1500*horasTrabajadas + (1500 * horasTrabajadas * 0.2)
	case "C":
		return 1000 * horasTrabajadas
	default:
		return 0
	}
}
