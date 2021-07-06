package main

import f "fmt"

func main() {
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

	var min float64
	var cat string
	f.Println("Ingrese los minutos trabajados")
	f.Scanln(&min)
	f.Println("Ingrese su categoria A, B o C")
	f.Scanln(&cat)
	f.Println("Su sueldo es de $", calculateSalary(min, cat))

}

func calculateSalary(min float64, cat string) float64 {
	var salary float64
	switch cat {
	case "C":
		salary = calcularHoras(min, 1000)
	case "B":
		mensual := calcularHoras(min, 1500)
		salary = mensual + ((mensual * 20) / 100)
	case "A":
		mensual := calcularHoras(min, 3000)
		salary = mensual + ((mensual * 50) / 100)
	}
	return salary
}

func calcularHoras(min float64, s float64) float64 {
	h := min / 60
	return h * s
}
