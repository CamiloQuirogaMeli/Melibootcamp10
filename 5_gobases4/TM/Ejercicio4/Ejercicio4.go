package main

import (
	"errors"
	"fmt"
)

func main() {

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
	var h int
	var s int
	var mt int
	var cant int
	fmt.Println("Ingrese las horas trabajadas")
	fmt.Scanln(&h)
	fmt.Println("Ingrese el pago por hora")
	fmt.Scanln(&s)
	fmt.Println("Ingrese mejor salario del semestre")
	fmt.Scanln(&mt)
	fmt.Println("Ingrese los meses trabajados en el semestre")
	fmt.Scanln(&cant)
	salary, err := salaryMonth(h, s)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Su salario mensual es de:", salary)
	}

	bonT, err := bonus(mt, cant)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Su aguinaldo es de:", bonT)
	}

}

func salaryMonth(h int, s int) (int, error) {
	salaryT := h * s
	if h < 80 || h < 0 {
		return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	} else if salaryT >= 150000 {
		return salaryT - (salaryT*10)/100, nil
	} else {
		return salaryT, nil
	}
}

func bonus(mt int, cantM int) (int, error) {
	bon := (mt / 12) * cantM
	if mt < 0 || cantM < 0 {
		return 0, errors.New("error: Se ingreso un valor negativo")
	} else {
		return bon, nil
	}
}
