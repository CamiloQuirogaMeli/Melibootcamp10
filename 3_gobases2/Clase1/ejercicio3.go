package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		Una empresa marinera necesita calcular el salario de sus empleados basándose en la
		cantidad de horas trabajadas por mes y la categoría.
		Si es categoría C, su salario es de $1.000 por hora
		Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
		Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

		Se solicita generar una función que se le pueda pasar como parámetros la cantidad de
		minutos trabajados por mes y la categoría, y nos devuelva su salario
	*/

	var entradoMinutos string
	var categoria string
	var minutos float64
	fmt.Println("Ingresa el número de minutos trabajados:")
	fmt.Scanln(&entradoMinutos)
	minutos, err := strconv.ParseFloat(entradoMinutos, 8)
	for err != nil {
		fmt.Println("El valor debe ser de tipo numérico")
		fmt.Println("Ingresa el número de minutos trabajados:")
		fmt.Scanln(&entradoMinutos)
		minutos, err = strconv.ParseFloat(entradoMinutos, 8)
	}
	fmt.Println("Ingresa la categoría del trabajador:")
	fmt.Scanln(&categoria)

	fmt.Println("El salario es: ", calcularSalario(minutos, categoria))

}

func calcularSalario(minutos float64, categoria string) float64 {
	var salario float64
	switch categoria {
	case "a", "A":
		salario = deMinAHoras(minutos) * 3000 * 1.5
	case "b", "B":
		salario = deMinAHoras(minutos) * 1500 * 1.2
	case "c", "C":
		salario = deMinAHoras(minutos) * 1000
	default:
		fmt.Println("Error, la categoría aún no existe")
	}

	return salario
}

func deMinAHoras(minutos float64) float64 {
	horas := minutos / 60
	return horas
}
