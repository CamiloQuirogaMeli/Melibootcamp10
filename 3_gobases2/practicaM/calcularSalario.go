package main

import "fmt"

func calcularSalario(minutos int, categoria string) int {
	var salario, hora, porcentaje int = 0, 1 / 60, 0
	hora *= minutos
	switch categoria {
	case "C":
		salario = 1000 * hora
	case "B":
		salario = 1500 * hora
		porcentaje = salario * 20 / 100
		salario += porcentaje
	case "A":
		salario = 3000 * hora
		porcentaje = salario * 50 / 100
		salario += porcentaje
	}

	return salario
}

func main() {

	fmt.Println("Salario total es: $", calcularSalario(120, "A"))

}
