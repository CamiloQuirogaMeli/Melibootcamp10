package main

import (
	"fmt"
)

type error interface {
	Error() string
}

// definimos un type struct
type customError struct {
	msg    string
	status int
}

func (e *customError) Error() string {
	return fmt.Sprintf("%d - %v", e.status, e.msg)
}

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
		return "", &customError{"error: el salario ingresado no alcanza el minimo imponible", 400}
	}

}
