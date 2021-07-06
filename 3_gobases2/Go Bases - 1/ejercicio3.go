package main

import "fmt"

/*
	Ejercicio 3 - Calcular salario
	Una empresa marinera necesita calcular el salario de sus empleados basándose en la
	cantidad de horas trabajadas por mes y la categoría.
	Si es categoría C, su salario es de $1.000 por hora
	Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
	Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

	Se solicita generar una función que se le pueda pasar como parámetros la cantidad de
	minutos trabajados por mes y la categoría, y nos devuelva su salario
*/

func calcularSalario(minutos int, categoría string) float64 {
	horas := float64(minutos) / 60
	salario := 0.0
	switch categoría {
	case "A":
		salario = horas*3000 + (horas * 3000 * 0.50)
	case "B":
		salario = horas*1500 + (horas * 1500 * 0.20)
	case "C":
		salario = horas * 1000
	default:
		fmt.Println("La categoria no corresponde a las existentes.")
	}
	return salario
}

func main() {
	minutos := 0
	categoria := ""
	fmt.Print("Digite los minutos del empleado trabajado por mes: ")
	fmt.Scan(&minutos)
	fmt.Print("Digite la categoria del salario mensual del empleado: ")
	fmt.Scan(&categoria)
	fmt.Println("El salario por mes del empleado es:", calcularSalario(minutos, categoria))
}
