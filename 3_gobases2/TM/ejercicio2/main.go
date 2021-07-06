package main

import (
	"errors"
	"fmt"
)

func main() {
	var option int
	var aux int = 1

	const banner string = "\n███╗   ███╗    ███████╗    ███╗   ██╗    ██╗   ██╗\n████╗ ████║    ██╔════╝    ████╗  ██║    ██║   ██║\n██╔████╔██║    █████╗      ██╔██╗ ██║    ██║   ██║\n██║╚██╔╝██║    ██╔══╝      ██║╚██╗██║    ██║   ██║\n██║ ╚═╝ ██║    ███████╗    ██║ ╚████║    ╚██████╔╝\n╚═╝     ╚═╝    ╚══════╝    ╚═╝  ╚═══╝     ╚═════╝ "

	const options string = "1. Calcular promedio de calificaciones \n 0. Salir"

	var nValues int
	for aux != 0 {
		fmt.Println(banner)
		fmt.Println(options)
		fmt.Scan(&option)

		switch option {
		case 1:
			fmt.Println("Ingrese el numero de calificaciones")
			fmt.Scan(&nValues)
			var prom float64 = scanValues(nValues)
			fmt.Println("*****PROMEDIO *****", "\n", prom)

		case 0:
			aux = 0
		default:
			fmt.Println("Opcion incorrecta!!!")
		}

	}

}

func scanValues(nValues int) (prom float64) {
	var total int
	for i := 0; i < nValues; i++ {
		var value int
		var aux_scan int = 1
		fmt.Println("Ingrese la califacion")
		fmt.Scan(&value)

		for aux_scan != 0 {
			aux_scan = 0
			hasError := evaluateValue(value)
			if hasError != nil {
				aux_scan = 1
				fmt.Println(hasError)
				fmt.Println("Ingrese la califacion")
				fmt.Scan(&value)
			}
		}
		total += value
	}
	return float64(total) / float64(nValues)
}

func evaluateValue(value int) error {
	if value >= 0 {
		return nil
	} else {
		return errors.New("¡¡ERROR!! has ingresado un valor negativo")
	}
}
