package main

import (
	"fmt"
)

func calcularSalario(salary int) (string, error) {
	if salary > 150000 {
		return "Debe pagar impuesto", nil
	} else {
		return "", fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %v", salary)
	}
}

func main() {
	/*

		Repite el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de
		error reciba por parámetro el valor de “salary” indicando que no alcanza el mínimo imponible
		(el mensaje mostrado por consola deberá decir: “error: el mínimo imponible es de 150.000 y
		el salario ingresado es de: [salary]”, siendo [salary] el valor de tipo int pasado por parámetro).

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
