package main

import "fmt"

func main() {
	fmt.Println("El salario del empleado es:", Ejercicio3(360, "A"))
}

func Ejercicio3(minutos uint, categoria string) (salario float64) {
	fmt.Println("Ejercicio3")
	const (
		ADICIONALCATEGORIAA float64 = 1.5
		ADICIONALCATEGORIAB float64 = 1.2
	)

	switch categoria {
	case "A":
		calculoSalarioBase := 3000.00 * (minutos / 60)
		salario = float64(calculoSalarioBase) * ADICIONALCATEGORIAA
	case "B":
		calculoSalarioBase := 1500.00 * (minutos / 60)
		salario = float64(calculoSalarioBase) * ADICIONALCATEGORIAB
	case "C":
		salario = float64(1000.00 * (minutos / 60))
	default:
		salario = 0.00
		fmt.Println("La categoria ingresada no existe")
	}
	return salario
}
