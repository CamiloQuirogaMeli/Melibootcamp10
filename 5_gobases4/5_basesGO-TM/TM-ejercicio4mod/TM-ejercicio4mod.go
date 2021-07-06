package main

import (
	"errors"
	"fmt"
	"os"
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

type error interface {
	Error() string
}

type errorTwo struct {
	msg string
}

func (e errorTwo) Error() string {
	return e.msg
}

func salarioMensual(hsTrabajadas int, valorHora float64) (float64, error) {

	if hsTrabajadas < 80 {
		return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}

	salario := valorHora * float64(hsTrabajadas)

	if salario >= 150000 {
		salario = salario * 0.9
	}

	return salario, nil

}

func medioAguinaldo(mejorSalario float64, mesesTrabajados int) (float64, error) {

	if mejorSalario < 0 {
		return 0, errorTwo{msg: "error: el salario no puede ser negativo"}
	}

	if mesesTrabajados < 0 {
		return 0, errorTwo{msg: "error: la cantidad de meses trabajados no puede ser negativo"}
	}

	meAguinaldo := mejorSalario / float64(12) * float64(mesesTrabajados)

	return meAguinaldo, nil
}

func main() {

	salario, err1 := salarioMensual(200, 800.0)

	if err1 != nil {
		fmt.Println(err1)
		os.Exit(1)
	}
	fmt.Println("El salario mensual es de: $", salario)

	medioAg, err2 := medioAguinaldo(-150000, 11)

	if err2 != nil {
		e1 := fmt.Errorf("e2: %w", err2)
		fmt.Println(errors.Unwrap(e1))
		fmt.Println(errors.Unwrap(err2))
		os.Exit(1)
	}

	fmt.Println("El medio aguinaldo es de: $", medioAg)

}
