package main

import (
	"errors"
	"fmt"
)

const HOUR = 60.0

var (
	category rune
	minutes  int
)

func main() {

	category = 'A'
	minutes = 2000

	salary, err := calculateSalary(minutes, category)

	if err == nil {
		fmt.Println(salary)
	} else {
		fmt.Println(err)
	}

}

func calculateSalary(min int, cat rune) (salary float64, err error) {

	err = nil

	if categoryValidation(cat) {
		salary = getSalary(min, cat)
		return
	} else {
		err = errors.New("categoria invalida")
		return
	}
}

func categoryValidation(cat rune) bool {
	return (cat == 'A' || cat == 'B' || cat == 'C')
}

func getSalary(min int, cat rune) float64 {

	switch cat {
	case 'A':
		return (((float64)(min) / HOUR) * 3000.0) + (3000.0 * 0.5)
	case 'B':
		return (((float64)(min) / HOUR) * 1500.0) + (3000.0 * 0.2)
	case 'C':
		return ((float64)(min) / HOUR) * 1000.0
	default:
		return 0.0
	}
}

/*
Una empresa marinera necesita calcular el salario de sus empleados basándose en la
cantidad de horas trabajadas por mes y la categoría.
Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

Se solicita generar una función que se le pueda pasar como parámetros la cantidad de
minutos trabajados por mes y la categoría, y nos devuelva su salario
*/
