package main

import (
	"fmt"
)

func main() {
	var option int
	var aux int = 1

	for aux != 0 {

		const banner string = "\n███╗   ███╗    ███████╗    ███╗   ██╗    ██╗   ██╗\n████╗ ████║    ██╔════╝    ████╗  ██║    ██║   ██║\n██╔████╔██║    █████╗      ██╔██╗ ██║    ██║   ██║\n██║╚██╔╝██║    ██╔══╝      ██║╚██╗██║    ██║   ██║\n██║ ╚═╝ ██║    ███████╗    ██║ ╚████║    ╚██████╔╝\n╚═╝     ╚═╝    ╚══════╝    ╚═╝  ╚═══╝     ╚═════╝ "

		const options string = "1. Impuestos de salario \n0. Salir"

		fmt.Println(banner)
		fmt.Println(options)
		fmt.Scan(&option)

		switch option {
		case 1:
			response, err := evaluateSalary()
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(response)

		case 0:
			aux = 0
		default:
			fmt.Println("Opcion incorrecta!!!")
		}
	}
}

func evaluateSalary() (string, error) {
	var salary int

	fmt.Println("Ingrese el valor del salario")
	fmt.Scan(&salary)

	if salary > 150000 {
		return "Debe pagar impuestos", nil
	} else {
		return "", fmt.Errorf("error: el minimio imponible es de 150000 y el salario ingresado es de: %d", salary)
	}

}
