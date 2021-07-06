package main

import (
	"errors"
	"fmt"
)

/*
	Ejercicio 4 - Impuestos de salario #4
	Vamos a hacer que nuestro programa sea un poco más complejo y útil.
	1. Desarrolla las funciones necesarias para permitir a la empresa calcular:

	a) Salario mensual de un trabajador según la cantidad de horas trabajadas. En caso de
		que el salario mensual sea igual o superior a $150.000, se le deberá descontar el 10%
		en concepto de impuesto. La función que se ocupe de realizar este cálculo deberá
		retornar más de un valor, incluyendo un error en caso de que la cantidad de horas
		mensuales ingresadas sea menor a 80 o un número negativo. El error deberá indicar
		“error: el trabajador no puede haber trabajado menos de 80 hs mensuales”.

	b) Calcular el medio aguinaldo correspondiente al trabajador (fórmula de cálculo de
		aguinaldo: [mejor salario del semestre] dividido 12 y multiplicar el [resultado
		obtenido] por la [cantidad de meses trabajados en el semestre]). La función que
		realice el cálculo deberá retornar más de un valor, incluyendo un error en caso de que
		se ingrese un número negativo.

	2. Desarrolla el código necesario para cumplir con las funcionalidades requeridas, utilizando
	“errors.New()”, “fmt.Errorf()” y “errors.Unwrap()”. No olvides realizar las validaciones de los
	retornos de error en tu función “main()”.
*/

func calcularSalario(horas int, valorHora int) (int, error) {
	salario := valorHora * horas

	if horas < 0 {
		return 0, errors.New("error: las horas ingresadas no pueden ser negativas")
	}

	if valorHora < 0 {
		return 0, fmt.Errorf("error: el valor ingresado es %d y debe ser mayor a 0", valorHora)
	}

	if horas < 80 {
		return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}

	if salario >= 150000 {
		salario -= (salario / 100 * 10)
		return salario, nil
	} else {
		salario += 0
		return salario, nil
	}
}

func calcularMedioAguinaldo(mesesTrabajados int, salarios ...int) (int, error) {
	mejorSalario := salarios[0]
	for _, ms := range salarios {

		if ms < 0 {
			return 0, errors.New("error: el salario debe ser un valor mayor a 0")
		}

		if mejorSalario > ms {
			mejorSalario = ms
		}
	}

	if mesesTrabajados < 0 {
		return 0, fmt.Errorf("error: la cantidad de meses trabajados es %d y debe ser mayor a 0", mesesTrabajados)
	}

	semestreAguinaldo := mejorSalario / 12 * mesesTrabajados
	return semestreAguinaldo, nil
}

func main() {
	var horas int
	var valorHora int
	var respuesta int

	fmt.Println("Ingrese la cantidad de las horas trabajadas al mes")
	fmt.Scan(&horas)
	fmt.Println("Ingrese el valor de la hora")
	fmt.Scan(&valorHora)

	salario, err := calcularSalario(horas, valorHora)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print("Desea calcular el aguinaldo 0.N0 1.SI: ")
		fmt.Scan(&respuesta)
		if respuesta == 0 {
			fmt.Println("El salario es: ", salario)
		} else {
			aguinaldo, err := calcularMedioAguinaldo(6, salario, salario+1000, salario+2000)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("El salario es: ", salario)
				fmt.Println("El aguinaldo del semestre actual es: ", aguinaldo)
				fmt.Println("El total del salario mensual + el aginaldo es: ", (salario + aguinaldo))
			}
		}
	}
}
