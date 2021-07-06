package main

// Una empresa marinera necesita calcular el salario de sus empleados basándose en la
// cantidad de horas trabajadas por mes y la categoría.
// Si es categoría C, su salario es de $1.000 por hora
// Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
// Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

// Se solicita generar una función que se le pueda pasar como parámetros la cantidad de
// minutos trabajados por mes y la categoría, y nos devuelva su salario
func calcSalary(minutes int, category string) float64 {
	var hours float64 = float64(minutes / 60)
	var salary float64
	switch category {
	case "A":
		salary = (hours * 3000) * 1.5
	case "B":
		salary = (hours * 1500) * 1.2
	case "C":
		salary = hours * 1000
	}

	return salary
}
