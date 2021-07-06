package main

import (
	"errors"
	"fmt"
)

func main() {
	var option int
	var aux int = 1

	for aux != 0 {

		const banner string = "\n███╗   ███╗    ███████╗    ███╗   ██╗    ██╗   ██╗\n████╗ ████║    ██╔════╝    ████╗  ██║    ██║   ██║\n██╔████╔██║    █████╗      ██╔██╗ ██║    ██║   ██║\n██║╚██╔╝██║    ██╔══╝      ██║╚██╗██║    ██║   ██║\n██║ ╚═╝ ██║    ███████╗    ██║ ╚████║    ╚██████╔╝\n╚═╝     ╚═╝    ╚══════╝    ╚═╝  ╚═══╝     ╚═════╝ "

		const options string = "1. Calcular salario mensual\n2. Calcular aguinaldo \n0. Salir"

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

		case 2:
			response, err := evaluateAguinaldo()
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
	var value int
	var hours int
	e1 := errors.New("error: el trabajador no puede haber trabajado menos de 80hs mensuales ")
	e2 := fmt.Errorf("error: ha ingresado valores negativos\n%w", e1)

	fmt.Println("Ingrese el valor de la hora trabajada")
	fmt.Scan(&value)

	if value < 0 {
		return "", fmt.Errorf("error: ha ingresado valores negativos")
	}
	fmt.Println("Ingrese el numero de horas trabajadas")
	fmt.Scan(&hours)

	if hours < 80 {
		if hours < 0 {
			return "", e2
		}
		return "", errors.Unwrap(e2)
	}

	total := value * hours

	switch {
	case total >= 150000:
		return fmt.Sprintf("El sueldo total del trabajador es: %.2f", float64(total)*0.9), nil
	case total >= 0:
		return fmt.Sprintf("El sueldo total del trabajador es:$ %d", total), nil

	default:
		return "", fmt.Errorf("error: ha ingresado un valor negativo")
	}

}

func evaluateAguinaldo() (string, error) {
	var bestSalary int
	var amountMonths int
	e1 := fmt.Errorf("error: ha ingresado valores negativos")

	fmt.Println("Ingrese el mejor salario del semestre")
	fmt.Scan(&bestSalary)
	if bestSalary < 0 {
		return "", e1
	}
	fmt.Println("Ingrese el numero de meses trabajados en el semestre")
	fmt.Scan(&amountMonths)
	if amountMonths < 0 {
		return "", e1
	}

	if amountMonths > 6 {
		return "", errors.New("error: el trabajador no puede tener mas de 6 meses trabajados en un semestre")
	}
	total := bestSalary / 12 * amountMonths

	return fmt.Sprintf("El aguinaldo del trabajador es igual a: %d", total), nil

}
