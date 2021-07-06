package main

import (
	"fmt"
)

func calculateSalary(category string, minutes ...int) float64 {
	result := 0.0
	for _, value := range minutes {
		minutesToHours := float64(value) / 60.00
		switch category {
		case "A":
			result += minutesToHours * 3000.00

		case "B":
			result += minutesToHours * 1500.00

		case "C":
			result += minutesToHours * 1000.00
		}
	}
	switch category {
	case "A":
		result = result * 1.5

	case "B":
		result = result * 1.2
	}
	return result
}

func main() {
	//Categoria C = 1k x hora
	//Categoria B = 1.5k x hora + 20%
	//Categoria A = 3k x hora + 50%
	//Funcion debe recibir un ellipsis con la cant de minutos trabajados y la categoria de
	//Retorna el salario mensual.

	finalSalary := calculateSalary("C", 50, 300, 40, 20, 90, 500, 200, 50)
	fmt.Printf("El salario de acuerdo a la categoria es: %.2f", finalSalary)
}
