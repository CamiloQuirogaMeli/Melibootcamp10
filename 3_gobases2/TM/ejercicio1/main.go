package main

import "fmt"

func main() {
	var option int
	var aux int = 1

	const banner string = "\n███╗   ███╗    ███████╗    ███╗   ██╗    ██╗   ██╗\n████╗ ████║    ██╔════╝    ████╗  ██║    ██║   ██║\n██╔████╔██║    █████╗      ██╔██╗ ██║    ██║   ██║\n██║╚██╔╝██║    ██╔══╝      ██║╚██╗██║    ██║   ██║\n██║ ╚═╝ ██║    ███████╗    ██║ ╚████║    ╚██████╔╝\n╚═╝     ╚═╝    ╚══════╝    ╚═╝  ╚═══╝     ╚═════╝ "

	const options string = "1. Evaluar impuesto de un salario \n 0. Salir"

	for aux != 0 {
		fmt.Println(banner)
		fmt.Println(options)
		fmt.Scan(&option)

		switch option {
		case 1:
			var salary int
			var taxes float64

			fmt.Println("Ingrese el salario")
			fmt.Scan(&salary)
			taxes = evaluateSalary(salary)

			fmt.Println("******** SALARIO ******", "\n", "$", salary, "\n ***** IMPUESTO ****", "\n", "$", taxes, "\n ****** TOTAL A PAGAR ****", "\n", "$", float64(salary)-taxes)
		case 0:
			aux = 0
		default:
			fmt.Println("Opcion incorrecta!!!")
		}

	}

}

func evaluateSalary(salary int) float64 {
	if salary > 150000 {
		return float64(salary) * 0.27
	} else if salary > 50000 {
		return float64(salary) * 0.17
	}

	return 0
}
