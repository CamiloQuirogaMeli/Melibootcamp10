package main

import "fmt"

func calculoSalario(cantMin int, categ byte) float32 {
	switch categ {
	case 'A':
		return float32(3000*(cantMin/60) + (50 * (3000 * (cantMin / 60)) / 100))
	case 'B':
		return float32(1500*(cantMin/60) + (20 * (1500 * (cantMin / 60)) / 100))
	case 'C':
		return float32(1000 * (cantMin / 60))
	default:
		return 0.0
	}
}

func main() {

	var horasTrab int
	var categ byte

	fmt.Printf("Ingrese la cantidad de horas trabajadas: ")
	fmt.Scanf("%d", &horasTrab)

	fmt.Printf("Ingrese la categoría del empleado: ")
	fmt.Scanf("%c", &categ)

	fmt.Printf("El salario del empleado es: %.2f", calculoSalario(horasTrab*60, categ))
}
