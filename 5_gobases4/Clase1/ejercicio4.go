package main

import (
	"errors"
	"fmt"
)

func CalcularSalarioMensual(horasTrabajados int, valorHoraTrabajada float64) (float64, error) {
	salario := float64(horasTrabajados) * valorHoraTrabajada
	if horasTrabajados < 80 {
		return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}
	if salario >= 150000 {
		fmt.Println("Se le deberá descontar el 10% en concepto de impuesto.")
		salario *= 0.90
	}
	return salario, nil
}

func CalcularMedioAguinaldo(mejorSalario float64, mesesTrabajados int) (float64, error) {
	aguinaldo := (mejorSalario / 12) * float64(mesesTrabajados)
	if mejorSalario <= 0 || mesesTrabajados <= 0 || mesesTrabajados > 6 {
		return 0, fmt.Errorf("El salario tiene que ser mayor a 0, usted digitó %v y debe haber trabajado entre 1 y 6 meses, usted trabajó %v", mejorSalario, mesesTrabajados)
	}

	return aguinaldo, nil
}

func main() {

	/**

		Vamos a hacer que nuestro programa sea un poco más complejo y útil.
	1. Desarrolla las funciones necesarias para permitir a la empresa calcular:
	a) Salario mensual de un trabajador según la cantidad de horas trabajadas. En caso de
	que el salario mensual sea igual o superior a $150.000, se le deberá descontar el 10%
	en concepto de impuesto. La función que se ocupe de realizar este cálculo deberá
	retornar más de un valor, incluyendo un error en caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo. El error deberá indicar
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

	var horasTrabajadas int
	var valorHoraTrabajada float64
	var mejorSalario float64
	var mesesTrabajados int

	salir := false
	var opcion int

	for !salir {
		fmt.Println("Digita una opción:")

		fmt.Println("\t 1: Calcular salario")
		fmt.Println("\t 2: Calcular medio aguinaldo")
		fmt.Println("\t 3: Salir")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			fmt.Println("Digita el numero de horas trabajadas")
			fmt.Scanln(&horasTrabajadas)
			fmt.Println("Digita el valor de la hora trabajada")
			fmt.Scanln(&valorHoraTrabajada)

			salario, err := CalcularSalarioMensual(horasTrabajadas, valorHoraTrabajada)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("El salario mensual es de", salario)
			}
		case 2:
			fmt.Println("Digita tu mejor salario en el semestre")
			fmt.Scanln(&mejorSalario)
			fmt.Println("Digita el número de meses trabajados")
			fmt.Scanln(&mesesTrabajados)

			aguinaldo, err := CalcularMedioAguinaldo(mejorSalario, mesesTrabajados)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("El medio aguinaldo es de", aguinaldo)
			}
		case 3:
			salir = true
		}
	}
}
