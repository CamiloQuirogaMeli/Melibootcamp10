package main

import (
	"fmt"
)

type myError struct {
	mensaje string
}

func (e *myError) Error() string {
	return fmt.Sprintf(e.mensaje)
}

func calcularSalario(salary int) (string, error) {
	if salary > 150000 {
		return "Debe pagar impuesto", nil
	} else {
		return "", &myError{
			mensaje: "error: el salario ingresado no alcanza el mínimo imponible",
		}
	}
}

func main() {

	/*

		En tu función “main”, define una variable llamada “salary” y asignale un valor de tipo
		“int”.
		2. Crea un error personalizado con un struct que implemente “Error()” con el mensaje
		“error: el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que
		“salary” sea menor a 150.000. Caso contrario, imprime por consola el mensaje “Debe
		pagar impuesto”.

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
