package main

import (
	"fmt"
	"errors"
)

//Vamos a hacer que nuestro programa sea un poco más complejo y útil.
//1. Desarrolla las funciones necesarias para permitir a la empresa calcular:
//a) Salario mensual de un trabajador según la cantidad de horas trabajadas. En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar el 10%
//en concepto de impuesto. La función que se ocupe de realizar este cálculo deberá retornar más de un valor, incluyendo un error en caso de que la cantidad de horas mensuales
//ingresadas sea menor a 80 o un número negativo. El error deberá indicar “error: el trabajador no puede haber trabajado menos de 80 hs mensuales”.
//b) Calcular el medio aguinaldo correspondiente al trabajador (fórmula de cálculo de aguinaldo: [mejor salario del semestre] dividido 12 y multiplicar el [resultado obtenido]
//por la [cantidad de meses trabajados en el semestre]). La función que realice el cálculo deberá retornar más de un valor, incluyendo un error en caso de que se ingrese un
//número negativo.
//2. Desarrolla el código necesario para cumplir con las funcionalidades requeridas, utilizando “errors.New()”, “fmt.Errorf()” y “errors.Unwrap()”. No olvides realizar las
//validaciones de los retornos de error en tu función “main()”.

type errNeg struct {}

func (e errNeg) Error() string {
	return "error: se ingresó un número negativo"
 }

func calculateSalary (horas int) (float64, error) {

	if horas < 0 {
		e1 := fmt.Errorf("error: %w", errNeg{})
		return 0, errors.Unwrap(e1)
	} else if horas < 80 {
		return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}

	salary := horas * 10000

	if salary >= 150000 {
		desc := float64(salary) * 0.1
		return float64(salary) - float64(desc), nil
	}

	return float64(salary), nil

}

func calculateAguin(salary float64, months int) (float64, error) {

	if salary < 0 || months < 0{
		e1 := fmt.Errorf("error: %w", errNeg{})
		return 0, errors.Unwrap(e1)
	}

	return (float64(salary) / 12) * float64(months), nil

}

func main() {

	horas := 80
	months := -6

	salary, err := calculateSalary(horas)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El salario mensual calculado es $", salary)

		aguinal, err := calculateAguin(salary, months)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("El aguinaldo es $", aguinal)
		}
	}

}