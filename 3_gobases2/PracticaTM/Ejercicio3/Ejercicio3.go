// Ejercicio 3 - Calcular salario

// Una empresa marinera necesita calcular el salario de sus empleados basándose en la
// cantidad de horas trabajadas por mes y la categoría.
// Si es categoría C, su salario es de $1.000 por hora
// Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
// Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

// Se solicita generar una función que se le pueda pasar como parámetros la cantidad de
// minutos trabajados por mes y la categoría, y nos devuelva su salario

package main

import (
	"fmt"
	"strings"
)

func calculateSalary(category string, minutesWork int) float64 {
	var salary float64
	category = strings.ToUpper(category)
	switch hoursWork := (float64(minutesWork) / 60); category {
	case "A":
		salaryPerHour := 3000.00
		salary = hoursWork * salaryPerHour
		salary += (salary * 0.5)
	case "B":
		salaryPerHour := 1500.00
		salary = hoursWork * salaryPerHour
		salary += (salary * 0.2)
	case "C":
		salaryPerHour := 1000.00
		salary = hoursWork * salaryPerHour
	}

	return salary
}

func main() {

	var category string
	var minutesWork int

	fmt.Printf("Ingrese la categoría de su empleado: ")
	fmt.Scan(&category)
	fmt.Printf("Ingrese los minutos trabajados de su empleado: ")
	fmt.Scan(&minutesWork)

	salary := calculateSalary(category, minutesWork)

	fmt.Printf("El salario mensual del empleado es: %.2f\n", salary)

}
