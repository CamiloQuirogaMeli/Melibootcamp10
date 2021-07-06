package main

import (
	"fmt"
)

func main() {
	var option int
	var aux int = 1
	for aux != 0 {

		const banner string = "\n███╗   ███╗    ███████╗    ███╗   ██╗    ██╗   ██╗\n████╗ ████║    ██╔════╝    ████╗  ██║    ██║   ██║\n██╔████╔██║    █████╗      ██╔██╗ ██║    ██║   ██║\n██║╚██╔╝██║    ██╔══╝      ██║╚██╗██║    ██║   ██║\n██║ ╚═╝ ██║    ███████╗    ██║ ╚████║    ╚██████╔╝\n╚═╝     ╚═╝    ╚══════╝    ╚═╝  ╚═══╝     ╚═════╝ "

		const options string = "1. Calcular salario del marinero \n 0. Salir"

		fmt.Println(banner)
		fmt.Println(options)
		fmt.Scan(&option)

		switch option {
		case 1:
			var minutes int
			var category string
			fmt.Println("Ingrese el numero total de minutos trabajados")
			fmt.Scan(&minutes)
			minutes = evaluateMinutes(minutes)

			fmt.Println("Ingrese la categoria")
			fmt.Scan(&category)

			category = evaluateCategory(category)
			salary := evaluateSalary(minutes, category)
			fmt.Println("***** SALARIO ****\n", "$", salary)

		case 0:
			aux = 0
		default:
			fmt.Println("Opcion incorrecta!!!")
		}
	}

}

func evaluateMinutes(minutes int) int {

	for minutes < 0 {
		fmt.Println("¡¡ERROR!! Ha ingresado un valor negativo")
		fmt.Println("Ingrese nuevamente el numero total de minutos trabajados")
		fmt.Scan(&minutes)
	}
	return minutes
}

func evaluateCategory(category string) string {

	for category != "a" && category != "A" &&
		category != "b" && category != "B" &&
		category != "c" && category != "C" {
		fmt.Println("¡¡ERROR!! Ha ingresado una categoria inexistente")
		fmt.Println("Ingrese nuevamente la categoria")
		fmt.Scan(&category)
	}
	return category
}

func evaluateSalary(minutes int, category string) float64 {
	var hours int = minutes / 60

	switch category {
	case "a", "A":
		var monthly int = 3000 * hours
		return float64(monthly) * 1.5
	case "b", "B":
		var monthly int = 1500 * hours
		return float64(monthly) * 1.2
	case "c", "C":
		var monthly int = 1000 * hours
		return float64(monthly)
	default:
		return 0.0
	}
}
