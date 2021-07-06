package main

import (
	"errors"
	"fmt"
)

const (
	MINIMO   = "minimo"
	PROMEDIO = "promedio"
	MAXIMO   = "maximo"
)

func main() {
	minFunc, err := operacion(MINIMO)
	promFunc, err := operacion(PROMEDIO)
	maxFunc, err := operacion(MAXIMO)
	if err != nil {
		fmt.Println(err)
	}

	min := minFunc(9, 9, 9, 9, 9, 9, 9, 9, 9, 2)
	prom := promFunc(2, 20)
	max := maxFunc(2, 2, 2, 2, 2, 2, 2, 2, 2, 5)

	fmt.Println("Minimo: ", min)
	fmt.Println("Promedio: ", prom)
	fmt.Println("Maximo: ", max)
}

func operacion(op string) (func(...int) int, error) {
	switch op {
	case MINIMO:
		return minimoFunc, nil
	case PROMEDIO:
		return promedio, nil
	case MAXIMO:
		return maxFunc, nil
	default:
		return nil, errors.New("Operación inválida")
	}
}

func minimoFunc(valores ...int) int {
	var minimo int

	for i, val := range valores {
		if i == 0 {
			minimo = val
			continue
		} else if val < minimo {
			minimo = val
		}
	}
	return minimo
}
func maxFunc(valores ...int) int {
	var minimo int

	for i, val := range valores {
		if i == 0 {
			minimo = val
		} else if val > minimo {
			minimo = val
		}
	}
	return minimo
}
func promedio(valores ...int) int {
	var suma int

	for _, val := range valores {
		suma += val
	}
	return suma / len(valores)
}
