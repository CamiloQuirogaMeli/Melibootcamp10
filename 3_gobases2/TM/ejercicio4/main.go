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

		const options string = "1. Iniciar \n 0. Salir"

		fmt.Println(banner)
		fmt.Println(options)
		fmt.Scan(&option)

		switch option {
		case 1:
			var action string
			fmt.Println("Ingrese la acciòn que desea realizar")
			fmt.Scan(&action)
			functionAction, err := evaluateAction(action)

			if err != nil {
				fmt.Println(err)
				break
			}
			var nValues int
			fmt.Println("Ingrese el numero de valores")
			fmt.Scan(&nValues)
			values := addValues(nValues)
			response := functionAction(values...)
			fmt.Println("*****RESPUESTA***\n", response)

		case 0:
			aux = 0
		default:
			fmt.Println("Opcion incorrecta!!!")
		}
	}

}

func evaluateAction(action string) (func(values ...int) string, error) {

	const (
		MINIMUM = "minimo"
		MAXIMUM = "maximo"
		AVERAGE = "promedio"
	)

	switch action {
	case MINIMUM:
		return doMinimum, nil
	case MAXIMUM:
		return doMaximum, nil
	case AVERAGE:
		return doAverage, nil
	default:
		return nil, errors.New("¡¡ERROR!! has ingresado una accion invalida")
	}
}

func addValues(nValues int) []int {
	var values []int

	for i := 0; i < nValues; i++ {
		var value int
		fmt.Println("Ingrese un valor")
		fmt.Scan(&value)
		values = append(values, value)
	}
	return values
}

func doMinimum(values ...int) string {
	var min int = 100000000000000000
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return "El valor minimo es " + fmt.Sprint(min)
}
func doMaximum(values ...int) string {
	var max int = 0
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return "El valor maximo es " + fmt.Sprint(max)
}
func doAverage(values ...int) string {
	var acc int = 0
	for _, value := range values {
		acc += value
	}
	return "El promedio es " + fmt.Sprint(acc/len(values))
}
