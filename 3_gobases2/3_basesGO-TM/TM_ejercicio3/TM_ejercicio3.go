package main

import (
	"errors"
	"fmt"
)

/*

Una empresa marinera necesita calcular el salario de sus empleados basándose en la
cantidad de horas trabajadas por mes y la categoría.
Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

Se solicita generar una función que se le pueda pasar como parámetros la cantidad de
minutos trabajados por mes y la categoría, y nos devuelva su salario

*/

/*Calculo el salario en función de la categoria*/

func calculaSalario(minutos int, categoria string) (float64, error) {

	salario := 0.0

	/*Chequeo la cantidad de minutos trabajados*/
	if minutos < 0 {
		return salario, errors.New("la cantidad de minutos trabajados no puede ser negativa")
	}

	/*Cantidad de horas mensuales trabajadas*/
	horas := float64(minutos / 60)

	/*Calculo salario según categoría */
	switch categoria {
	case "A":
		salario = 3000 * horas * 1.5
	case "B":
		salario = 1500 * horas * 1.2
	case "C":
		salario = 1000 * horas
	default:
		return salario, errors.New("la categoria no es válida")
	}

	return salario, nil
}

func main() {

	minutos := 8 * 60 * 20
	categoria := "A"

	salario, err := calculaSalario(minutos, categoria)

	fmt.Println("Detalle de salario")
	fmt.Println("Minutos trabajados: ", minutos)
	fmt.Println("Categoria: ", categoria)
	fmt.Println("Salario: ", salario)
	fmt.Println("Error en la operación: ", err)

}
