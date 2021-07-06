package main

import (
	"errors"
	"fmt"
)

func calcularSalario(salary int) (string, error) {
	if salary > 150000 {
		return "Debe pagar impuesto", nil
	} else {
		return "", errors.New("error: el salario ingresado no alcanza el mínimo imponible")

	}
}

func main() {

	/*

		Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en
		reemplazo de “Error()”, se implemente “errors.New()”.

	*/

	var salary int

	salir := false
	var opcion int

	for !salir {
		fmt.Println("Digita una opción:")

		fmt.Println("\t 1: Digitar salario")
		fmt.Println("\t 2: Salir")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			fmt.Println("Digita el valor del salario")
			fmt.Scanln(&salary)
			mensaje, err := calcularSalario(salary)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(mensaje)
			}
		case 2:
			salir = true
		}
	}
}
