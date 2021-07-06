/* Ejercicio 4 - Impuestos de salario #4
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
retornos de error en tu función “main()”. */
package main

import (
	"fmt"
	"errors"
)

const valorHora = 1000

func main() {
	salario, impuesto, err := calcularSalario(50)
	if err == nil {
		fmt.Printf("Salario: %d - Se descontaron impuestos: %v\n", salario, impuesto)
	} else {
		fmt.Println(fmt.Errorf("%s", err.Error()))
	}
}

func calcularSalario(horas int) (int, bool, error) {
	if horas < 80 {
		return 0, false, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	} else {
		val := horas * valorHora
		if val < 150000 {
			return val, false, nil
		}
		return val - (10 * val / 100), true, nil
	}
}